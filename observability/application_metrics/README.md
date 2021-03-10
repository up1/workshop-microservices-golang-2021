# Step to run

## 1. Start service
```
$go mod tidy
$go run cmd/main.go
```

## 2. Testing
URL for testing
* Metric endpoint :: http://localhost:8080/metrics
* APIs
  * Get all beers :: http://localhost:8080/beer/
  * Get beer by id
    * Found => http://localhost:8080/beer/1
    * Not found => http://localhost:8080/beer/0
