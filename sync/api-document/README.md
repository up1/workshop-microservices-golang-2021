# Step to run

## 1. Generate api document with Swagger
* Using [swag](https://github.com/swaggo/swag)

```
$swag init --dir cmd/ --parseDependency --output docs

2021/03/09 19:58:53 Generate swagger docs....
2021/03/09 19:58:53 Generate general API Info, search dir:cmd/
2021/03/09 19:58:53 Generating beer.BeerResponse
2021/03/09 19:58:53 create docs.go at docs/docs.go
2021/03/09 19:58:53 create swagger.json at docs/swagger.json
2021/03/09 19:58:53 create swagger.yaml at docs/swagger.yaml
```

## 2. Start service
```
$go mod tidy
$go run cmd/main.go
```

URL for testing
* [API Documentation](http://localhost:8080/docs/index.html)
* [List of all beers](http://localhost:8080/beer/)