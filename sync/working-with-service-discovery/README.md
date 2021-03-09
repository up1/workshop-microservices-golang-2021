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
$docker-compose build
$docker-compose up -d
$docker-compose ps
$docker-compose logs --follow
```

