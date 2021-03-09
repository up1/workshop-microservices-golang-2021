# Step to run

## 1. Logging with [Logrus library](https://github.com/sirupsen/logrus) (JSON Formatter)
```
$go mod tidy
$go run main.go
```

## 2. Create wrapper log for Logrus library (JSON Formatter)
```
$go run main_better.go
```

## 3. Suggestion for logging
* JSON Format (Easy to use for centralize logging system)
* Use HTTP headers and unique IDs to log user behavior across microservices

## 4. Add UUID to logging (Working with distributed tracing)
* [Tracing message format](https://opentelemetry.io/)
* Distribited tracing server
  * [Jaeger](https://www.jaegertracing.io/)
  * [Zipkin](https://zipkin.io/)
