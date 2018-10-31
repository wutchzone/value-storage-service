# Value storage service

# Routes

## Save new value
```
POST /api/data/{:unit:}

{
    "value": "FLOAT"
}
```

## Get record list

```
GET /api/data/{:unit:}
```

## Get single record from list
```
GET /api/data/{:unit}/{:uuid:}
```