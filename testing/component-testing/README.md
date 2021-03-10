# Testing with Database

## 1. Start PostgreSQL database server
```
$cd postgresql
$docker-compose up -d
$docker-compose ps
```

Check data in postgresql database server
```
$docker container exec -it postgres bash

// Connect to database server
>psql -U user -W demo

// List all tables
>\dt

       List of relations
 Schema | Name  | Type  | Owner
--------+-------+-------+-------
 public | users | table | user
(1 row)

// Query data
>select * from users;

 id | name  | age
----+-------+-----
  1 | name1 |  20
  2 | name2 |  30
  3 | name3 |  40
(3 rows)
```

## 2. Start web server

```
$export POSTGRES_URL=postgres://user:password@localhost/demo?sslmode=disable 
$go mod tidy
$go run cmd/main.go
```

Access to API with URL `http://localhost:8080/users`

## 3. Testing with [Postman](https://www.postman.com/) + [newman cli](https://www.npmjs.com/package/newman)
* [Working with JSON SChema](https://json-schema.org/understanding-json-schema/)
```
// Install newman via npm
$npm install -g newman

// Testing with newman
$cd postman
$newman run *.json
```

## 4. Automated process with Docker
* Build Docker image of web server
* Start postgresql database server
* Start web server
* API testing with Postman


# Testing with External services

## 1. Start web server

```
$export API_URL=https://jsonplaceholder.typicode.com
$go mod tidy
$go run cmd/main.go
```

Access to API with URL `http://localhost:8080/users/external`

## 2. Testing with Postman
```
$cd postman
$newman run *.json
```

## 3. Stub/Mock External services

List of tools
* [Stubby4j](https://github.com/azagniotov/stubby4j)
* [Stubby4node](https://github.com/mrak/stubby4node)
* [JSON Server](https://github.com/typicode/json-server)
* [WireMock](http://wiremock.org/)
* [MounteBank](http://www.mbtest.org/)

### Example with Stubby4J
```
$cd mock_api
$java -jar java -jar stubby4j-x.x.xx.jar -d api.yml
```


Docker
```
$cd mock_api
$docker container run -d --rm \
    -e YAML_CONFIG=api.yml \
    -v $(pwd):/home/stubby4j/data \
    -p 8882:8882 -p 8889:8889 -p 7443:7443 \
    azagniotov/stubby4j:latest-jre8
```

Call mock api = http://localhost:8882/users


## 4. Stop and Start web server

```
$export API_URL=http://localhost:8882
$go mod tidy
$go run cmd/main.go
```

## 2. Testing with Postman
```
$cd postman
$newman run *.json
```
