## Compose Simple Go + Nginx + MSSQL Server DB

### Nginx Proxy with Go backend

Project structure:
```
.
|-- backend
|   |-- Dockerfile
|   |-- main.go
|   |-- go.mod
|   |-- go.sum
|-- proxy
|   |-- nginx.conf
|-- sql
|   |-- data
|   |   |-- ...
|   |-- init.sql
|-- README.md
|-- compose.yml
|-- .env
```

## Deploy with docker compose

```
$ docker compose up -d
PASTE IN CMD RESULT
```

## Expected result

Listing the containers must show 3 containers running and the port mapping as below:
```
$ docker compose ps
PASTE IN CMD RESULT
```

After the application starts, navigate to `http://localhost:8000/order_status/1000` in your browser or run:
```
$ curl localhost:8000/order_status/1000
PASTE IN CMD RESULT
```

Stop and remove the containers
```
$ docker compose down
```
