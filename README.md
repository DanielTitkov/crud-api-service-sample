# crud-api-service-sample

Sample Golang JSON API service built with Gin and Entgo.

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