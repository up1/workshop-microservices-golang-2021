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



