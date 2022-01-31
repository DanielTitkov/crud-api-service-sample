# crud-api-service-sample

Sample Golang JSON API service built with Gin and Entgo. 

## 🍕🍕🍕

## Requirements 

In order to run service the following is required:
* Golang (tested with 1.16)
* Make (in order to use Makefile commands)
* Docker (to run database in container)

## Commands

* `make run` to run service locally
* `make test` to run tests
* `make lint` to run linters (golangci-lint)
* `make check` to run both tests and linters
* `make build` to build service binary
* `make db` to start dev database (with docker)
* `make entgen` regenerate entgo files

## Configuration 

Configuration is passed as an argument e.g. `go run cmd/app/main.go ./configs/dev.yml`. Database URI, environment and server parameters can be specified in this file. 

## Methods

Standart responses and empty requests are not shown here. 

### /api/v1/createPizza

**Request**

```json
{
    "title": "Fiery",
    "price": 77700,
    "dougn": "thick",
    "description": "as spicy as expensive"
}
```

### /api/v1/updatePizza

Title is immutable and can not be updated.

**Request**

```json
{
    "id": 13,
    "price": 600,
    "dough": "thin",
    "description": "Now chean and thin"
}
```

### /api/v1/getPizzas

**Response**

```json
{
    "pizzas": [
        {
            "id": 1,
            "title": "Fiery",
            "price": 77700,
            "description": "as spicy as expensive",
            "dough": "thick"
        },
        {
            "id": 2,
            "title": "Carbonara",
            "price": 1000,
            "description": "best pizza",
            "dough": "thin"
        }
    ]
}
```

### /api/v1/getPizzaByID

**Request**

```json
{
    "id": 1
}
```

**Response**

```json
{
    "id": 1,
    "title": "Fiery",
    "price": 77700,
    "description": "as spicy as expensive",
    "dough": "thick"
}
```

### /api/v1/deletePizzaByID

**Request**

```json
{
    "id": 9
}
```