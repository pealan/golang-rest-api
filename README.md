# golang-rest-api

REST service that supports the management of a device database.

### Prerequisites
1. Docker
2. Make

## Build, run and test

```bash
docker compose up --build
```

Once you see this message in the terminal, it means the server is up!
```bash
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```

### Endpoints and examples:

#### Swagger documentation: http://localhost:8080/swagger/index.html

#### Add device

```bash
curl http://localhost:8080/device --include --header "Content-Type: application/json" --request "POST" --data '{"name": "iPhone 11", "brand": "Apple"}'
```

#### List all devices

```bash
curl http://localhost:8080/device
```

#### List all devices with specific brand name

```bash
curl http://localhost:8080/device?brand=Apple
```

#### Get device by ID

```bash
curl http://localhost:8080/device/1
```

#### Update device by ID (brand and/or name)

```bash
curl http://localhost:8080/device/1 --include --header "Content-Type: application/json" --request "PUT" --data '{"name": "iPhone 15X", "brand": "Apple"}'

# or partially

curl http://localhost:8080/device/1 --include --header "Content-Type: application/json" --request "PUT" --data '{"name": "iPhone 15X"}'
```

#### Delete device by ID

```bash
curl http://localhost:8080/device/1 --include --request "DELETE"
```

## Unit tests

```bash
make test
```

Check terminal for test information.

## Development tools

### Environment
1. Golang 1.23+
1. Docker
1. Make
1. Debian-based distro (recommended)

### Wire

After adding new endpoints to this API, you will need to run Wire to inject the new dependencies to the server. To do so, inside the root directory, run:
```bash
go run github.com/google/wire/cmd/wire@latest
```
It'll (re)generate `wire_gen.go`, which will automatically implement the dependency injections. For more info, check the [official repository](https://github.com/google/wire).

### Swagger

Read the [Golang Swagger documentation](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format) to check expected attributes. After adding or updating documentation for any route handlers, run:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init
```
It'll (re)generate files under the `docs` directory and update Swagger documentation once the server is up. For more info, check the [official repository](https://github.com/swaggo/gin-swagger).

### GORM

We are using this package as ORM to connect to the application's database. Check [official documentation](https://gorm.io/docs/) for more info.

### Gin

We are using Gin as our server web framework. It simplifies routing and allows for an easier future middleware implementation. Check [official documentation](https://gin-gonic.com/docs/) for more info.

## Current issues and TODO's

1. Implement unit testing for the server handlers
1. Implement configuration file reading. It can be used to define database connection attributes, port usage, different database choices, etc.
1. Implement unified server response structure, useful if more endpoint groups are added.
1. Improve/Clean server logging
1. Implement custom error responses and project's standard
1. Configure DEBUG/PROD Gin settings

## Other tips

### Bring service up (no build required)

```bash
docker compose up
```

### Bring services down

```bash
docker compose down
```

### Reset database

```bash
docker compose down # if server is up
docker volume rm golang-rest-api_db-data
docker compose up --build # bring server and database up again
```


