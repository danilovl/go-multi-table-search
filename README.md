# GO simple multi table search #

## About ##

Simple multi table search where one query can search in several tables by certain parameters.

### Requirements

* Golang 1.18
* SQL DB

### 1. Installation

Change `DATABASE_DSN` in `.env` or create `.env.local`

``` text
DATABASE_DSN="root:root@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
```

Run command:

``` bash
go run main.go
```

### 2. Usage

Create POST request to `localhost:8080/multi-table-search` with `json`

``` json
[
  {
    "identifier": "product-1",
    "tableName": "product",
    "search": "orb",
    "selectColumns": ["id", "name", "created_at"],
    "whereColumns": ["name"],
    "limit": 100
  },
  {
    "identifier": "product-2",
    "tableName": "product",
    "search": "voda",
    "selectColumns": ["id", "name", "created_at"],
    "whereColumns": ["name"],
    "limit": 100
  },
  {
    "identifier": "product-3",
    "tableName": "product",
    "search": "m&m",
    "selectColumns": ["id", "name", "created_at"],
    "whereColumns": ["name"],
    "limit": 100
  },
  {
    "identifier": "product-4",
    "tableName": "product",
    "search": "baget",
    "selectColumns": ["id", "name", "created_at"],
    "whereColumns": ["name"],
    "limit": 100
  },
  {
    "identifier": "product-5",
    "tableName": "product",
    "search": "a",
    "selectColumns": ["id", "name", "created_at"],
    "whereColumns": ["name"],
    "limit": 100
  },
  {
    "identifier": "city-1",
    "tableName": "city",
    "selectColumns": ["*"],
    "limit": 100,
    "offset": 2
  }
]
```
