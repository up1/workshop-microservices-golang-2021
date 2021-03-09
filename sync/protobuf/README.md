# Step to run

## 1. Install [protoc](https://grpc.io/docs/protoc-installation/)

## 2. Generate the golang code from the Proto definition files

Install [command line tools](https://grpc.io/docs/languages/go/quickstart/)
```
$go get google.golang.org/protobuf/cmd/protoc-gen-go
$go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Generate golang code with protoc
```
$protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto
```

## 3. Start GRPC server
```
$go run cmd/server.go
```

## 4. Call GRPC server from client
```
$go run cmd/client.go
```

## 5. Benchmark GRPC server with [ghz](https://ghz.sh/)
```
$ghz --insecure --proto ./proto/user.proto --call users.Users.GetUsers 0.0.0.0:50000

// Config number of requests and concurrency users
$ghz --insecure --proto ./proto/user.proto --call users.Users.GetUsers -n 10000 -c 100 0.0.0.0:50000

Summary:
  Count:	10000
  Total:	225.41 ms
  Slowest:	7.19 ms
  Fastest:	0.14 ms
  Average:	1.71 ms
  Requests/sec:	44364.53

Response time histogram:
  0.140 [1]	|
  0.845 [1347]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.550 [3102]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.255 [3184]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.960 [1841]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.665 [393]	|∎∎∎∎∎
  4.370 [76]	|∎
  5.075 [27]	|
  5.780 [12]	|
  6.485 [5]	|
  7.191 [12]	|

Latency distribution:
  10 % in 0.77 ms
  25 % in 1.07 ms
  50 % in 1.69 ms
  75 % in 2.22 ms
  90 % in 2.63 ms
  95 % in 2.98 ms
  99 % in 3.86 ms

Status code distribution:
  [OK]   10000 responses
```


## References
* [Awesome GRPC](https://github.com/grpc-ecosystem/awesome-grpc)
* [gRPC benchmarking and load testing](https://grpc.io/docs/guides/benchmarking/)