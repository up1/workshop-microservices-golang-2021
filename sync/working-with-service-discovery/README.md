# Step to run

## 1. Download [consul](https://www.consul.io/)
Start consul service
```
$consul
```
Open consul ui `http://localhost:8500`

## 2. Start product service
```
$cd product-service
$go run main.go
```

Open in browser `http://localhost:8100/products`

## 3. Start user service
```
$cd user-service
$go run main.go
```

Open in browser `http://localhost:8080/user-products`

## 4. Working with Docker
```
$cd deploy

// Build Docker image
$docker-compose build

// Start consul service
$docker-compose up -d consul
$docker-compose ps

// Start all services
$docker-compose up -d
$docker-compose ps

// See all logs from all services
$docker-compose logs --follow
```

