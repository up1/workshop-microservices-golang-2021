# Step to run

## 1. Start with manual process

Step 1 :: Create [Jaeger tracing server](https://www.jaegertracing.io/)
```
$docker container run \
	-d --rm --name tracer \
	-p 6831:6831/udp \
	-p 6832:6832/udp \
	-p 16686:16686 \
	jaegertracing/all-in-one \
	--log-level=debug
```
Open in browser url=`http://localhost:16686/`

Step 2 :: Start all services
```
$go run cmd/service1/main.go
$go run cmd/service2/main.go
$go run cmd/service3/main.go
```

Step 3 :: Testing 
* Call service 1 with `http://localhost:8000/call`
* Call service 1 -> service 2 with `http://localhost:8000/call-service2`
* Call service 1 -> service 2 -> service 3 with `http://localhost:8000/call-service3`

## 2. Start with make file
```
$make bye
$make love
```
