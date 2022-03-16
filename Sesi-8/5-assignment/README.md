# Orders API


### How to run 
- clone this project
- change directory to Sesi-8/5-assignment
- create database example in this project (Sesi-8/5-assignment/db.sql)
- go run main.go

### GET Order API

This API is used to get orders data from Database

```
curl --request GET 'curl --location --request GET 'http://localhost:8080/order/{orderID}'
```

### Response

```
{
  "data": {
    "order_id": 6,
    "customer_name": "Khairul Abdi Dongoran",
    "ordered_at": "2019-11-09T21:21:46Z",
    "items": [
      {
        "item_id": 31,
        "item_code": "ABC12",
        "description": "A random description new",
        "quantity": 10,
        "order_id": 6
      },
      {
        "item_id": 32,
        "item_code": "ABC13",
        "description": "A random description new",
        "quantity": 100,
        "order_id": 6
      }
    ]
  }
}
```

### GET Orders API

This API is used to get orders data from Database

```
curl --request GET 'curl --location --request GET 'http://localhost:8080/orders'
```

### Response

```
{
  "data": [
    {
      "order_id": 1,
      "customer_name": "Khairul",
      "ordered_at": "2019-11-09T00:00:00Z",
      "items": [
        {
          "item_id": 14,
          "item_code": "ABC12",
          "description": "A random description new",
          "quantity": 10,
          "order_id": 1
        },
        {
          "item_id": 15,
          "item_code": "ABC13",
          "description": "A random description new",
          "quantity": 100,
          "order_id": 1
        }
      ]
    },
    {
      "order_id": 3,
      "customer_name": "Diah",
      "ordered_at": "2019-11-09T00:00:00Z",
      "items": [
        {
          "item_id": 18,
          "item_code": "A1B2C3",
          "description": "A random description",
          "quantity": 1,
          "order_id": 3
        },
        {
          "item_id": 19,
          "item_code": "A1B2C3",
          "description": "A random description",
          "quantity": 1,
          "order_id": 3
        }
      ]
    }
  ]
}
```

### GET Create Order API

This API is used to create order data to Database

```
curl --location --request POST 'http://localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
  "customer_name": "Khairul",
  "ordered_at": "2019-11-09T21:21:46+00:00",
  "items": [
    {
      "item_code":"A1B2C3",
      "description": "A random description",
      "quantity": 1
    },
    {
      "item_code":"A1B2C3",
      "description": "A random description",
      "quantity": 1
    }
  ]
}'
```

### Response

```
{
  "data": {
    "customer_name": "Khairul",
    "ordered_at": "2019-11-09T21:21:46Z",
    "items": [
      {
        "item_code": "A1B2C3",
        "description": "A random description",
        "quantity": 1
      },
      {
        "item_code": "A1B2C3",
        "description": "A random description",
        "quantity": 1
      }
    ]
  }
}
```

### GET Update Order API

This API is used to update order data to Database

```
curl --location --request PUT 'http://localhost:8080/orders/6' \
--header 'Content-Type: application/json' \
--data-raw '{
  "customer_name": "Khairul Abdi Dongoran",
  "ordered_at": "2019-11-09T21:21:46+00:00",
  "items": [
    {
      "item_code":"ABC12",
      "description": "A random description new",
      "quantity": 10
    },
    {
      "item_code":"ABC13",
      "description": "A random description new",
      "quantity": 100
    }
  ]
}'
```

### Response

```
{
  "data": {
    "customer_name": "Khairul Abdi Dongoran",
    "ordered_at": "2019-11-09T21:21:46Z",
    "items": [
      {
        "item_code": "ABC12",
        "description": "A random description new",
        "quantity": 10
      },
      {
        "item_code": "ABC13",
        "description": "A random description new",
        "quantity": 100
      }
    ]
  }
}
```


### GET Delete Order API

This API is used to delete order data from Database

```
curl --location --request DELETE 'http://localhost:8080/orders/2'
```

### Response

```
{
  "Message": "Deleted successfully"
}
```
