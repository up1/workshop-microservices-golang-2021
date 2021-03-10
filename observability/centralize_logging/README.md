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

## 5. Working with ELK stack
* Elasticsearch
* Logstash but Use [Filebeat](https://www.elastic.co/guide/en/beats/filebeat/current/index.html)
* Kibana

### Step 1 :: Start Elasticsearch
```
$cd elasticsearch-7.11.2
$bin/elasticsearch
```

Check status of Elasticseach server at http://localhost:9200/

### Step 2 :: Start Kibana
```
$cd kibana-7.11.2-darwin-x86_64
$bin/kibana
```

Check status of Kibana server at http://localhost:5601/

### Step 3 :: Working with File beats
* Monitoring data from log files
* Send data to Elasticsearch server

```
$filebeat -e -c filebeat.yml
```