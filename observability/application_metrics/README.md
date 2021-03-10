# Step to run

## 1. Start service
```
$go mod tidy
$go run cmd/main.go
```

## 2. Testing
URL for testing
* Metrics page for Prometheus => http://localhost:8080/metrics
* List of all beers => http://localhost:8080/beer
* Get beer by id (found) => http://localhost:8080/beer/1
* Get beer by id (not found) => http://localhost:8080/beer/1

## Download and start [Prometheus](https://prometheus.io/) and [Grafana](https://grafana.com/) 
Server
* Prometheus = Database server
  * port=9090
* Grafana = Visualization server
  * port=3000

Start Prometheus
```
$docker container run -p 9090:9090 -v $(pwd)/prometheus-data:/prometheus-data prom/prometheus --config.file=/prometheus-data/prometheus.yml
```
Open url=http://localhost:9090 in browser

Start Grafana
```
$docker container run -d -i -p 3000:3000 grafana/grafana
```

Open url=http://localhost:3000 in browser
* Username=admin
* Password=admin

Add more dashboard for [Go Gin Prometheus](https://github.com/zsais/go-gin-prometheus)

