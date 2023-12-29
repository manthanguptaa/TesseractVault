# Tesseract Vault
A simple in-memory key-value store server, implemented in Go.

## Features

- Stores key-value pairs of strings.
- Provides a RESTful API for setting and getting values.
- Uses interfaces for flexibility and testability.
- Employs mutexes for concurrent access.

## Getting Started

1. Install dependencies:
```bash
go get github.com/labstack/echo/v4
go get github.com/stretchr/testify/assert
```

2. Run the server:
```bash
go run main.go
```

3. Access the API:
- Set a value: ```curl http://localhost:3000/set/mykey/myvalue```
- Get a value: ```curl http://localhost:3000/get/mykey```
