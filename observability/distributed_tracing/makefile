SHELL := /bin/bash

love: tracer start

start:
	go build demo/cmd/service1
	go build demo/cmd/service2
	go build demo/cmd/service3
	./service1 &
	./service2 &
	./service3 &

tracer:
	docker container run \
	-d --rm --name tracer \
	-p 6831:6831/udp \
	-p 6832:6832/udp \
	-p 16686:16686 \
	jaegertracing/all-in-one \
	--log-level=debug

bye:
	pkill service1 || true
	pkill service2 || true
	pkill service3 || true
	docker container stop tracer