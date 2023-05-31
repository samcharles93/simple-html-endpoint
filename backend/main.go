package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		dbHost, dbUser, dbPass, dbPort, dbName)

	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/order_status/:id", getOrderStatus)

	router.Run(":8000")
}

func getOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	status, err := fetchOrderStatusFromDB(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Failed to retrieve the order status for ID %s", orderID),
		})
		return
	}

	canBuild := false
	reason := "Order not submitted"

	if status == "Sent for Provisioning" || status == "Complete" {
		canBuild = true
		reason = "Ok to build"
	}

	c.JSON(http.StatusOK, gin.H{
		"order_id":  orderID,
		"can_build": canBuild,
		"reason":    reason,
	})
}

func fetchOrderStatusFromDB(orderID string) (string, error) {
	query := "SELECT OrderStatus FROM Orders WHERE OrderID = ?"

	row := db.QueryRow(query, orderID)

	var orderStatus string
	err := row.Scan(&orderStatus)
	if err != nil {
		return "", err
	}

	return orderStatus, nil
}
