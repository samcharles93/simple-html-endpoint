USE master;

CREATE DATABASE CloudCloudly;
GO

USE CloudCloudly;

CREATE TABLE Orders (
  OrderID INT PRIMARY KEY,
  CustomerName VARCHAR(255) NOT NULL,
  OrderDate DATETIME2 NOT NULL,
  Product VARCHAR(255),
  Quantity INT,
  Price DECIMAL(10, 2),
  OrderStatus VARCHAR(50) CHECK (OrderStatus IN ('Sent for Provisioning', 'Complete', 'Cancelled', 'Draft', 'Quoted'))
);
GO

DECLARE @i int = 1;
DECLARE @status VARCHAR(50);
DECLARE @orderID INT;
WHILE @i <= 25
BEGIN
  SET @status = CASE (@i % 5)
    WHEN 0 THEN 'Sent for Provisioning'
    WHEN 1 THEN 'Complete'
    WHEN 2 THEN 'Draft'
    WHEN 3 THEN 'Quoted'
    WHEN 4 THEN 'Cancelled'
    ELSE 'Draft'
  END;

  SET @orderID = CAST((RAND() * (9999 - 1000 + 1)) + 1000 AS INT);

  INSERT INTO Orders(OrderID, CustomerName, OrderDate, Product, Quantity, Price, OrderStatus)
  VALUES (@orderID, 'Customer' + CAST(@i AS VARCHAR(10)), GETDATE(), 'Product' + CAST(@i AS VARCHAR(10)), @i, @i * 10.00, @status);
  SET @i = @i + 1;
END
GO
