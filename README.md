# Test golang service

## Reindexer DB

- `cd ./reindexer`
- Run `docker-compose -f .\docker-compose.yml up`

## Document service

- Run `go run cmd/testApp/main.go`

### API

- POST `/addDocument`

Body:
```
{
  "id": 2,
  "description": "DocumentDescription",
  "options": [
    {
      "id": 118,
      "sort": 25,
      "description": "DescriptionONew",
      "values": [
        { "id": 81, "title": "TitleNew" },
        { "id": 81, "title": "TitleNew" }
      ]
    },
    {
      "id": 162,
      "sort": 89,
      "description": "DescriptionONew",
      "values": [
        { "id": 40, "title": "TitleNew" },
        { "id": 11, "title": "TitleNew" }
      ]
    },
    {
      "id": 31,
      "sort": 29,
      "description": "DescriptionONew",
      "values": [
        { "id": 90, "title": "TitleNew" },
        { "id": 87, "title": "TitleNew" }
      ]
    }
  ]
}
```

- POST `/getAllDocuments`

Body:
```
{
	"skip": 0, // optional
	"take": 2 // optional
}
```

- POST `/getDocumentById`

Body:
```
{
	"documentId": 2
}
```

- POST `/updateDocument`

Body:
```
{
  "id": 2,
  "description": "DocumentDescription",
  "options": [
    {
      "id": 118,
      "sort": 25,
      "description": "DescriptionONew",
      "values": [
        { "id": 81, "title": "TitleNew" },
        { "id": 81, "title": "TitleNew" }
      ]
    },
    {
      "id": 162,
      "sort": 89,
      "description": "DescriptionONew",
      "values": [
        { "id": 40, "title": "TitleNew" },
        { "id": 11, "title": "TitleNew" }
      ]
    },
    {
      "id": 31,
      "sort": 29,
      "description": "DescriptionONew",
      "values": [
        { "id": 90, "title": "TitleNew" },
        { "id": 87, "title": "TitleNew" }
      ]
    }
  ]
}
```

- POST `/deleteDocumentById`

Body:
```
{
	"documentId": 2
}
```

To test create/update API you can use [test client app](https://github.com/BMO568/golang-test-service/blob/master/client/httpClient.go).
