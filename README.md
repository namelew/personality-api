# REST API IN GOLANG
A personalities registrer api, using golang in backend with gorilla as router and handler and gorm as database interface
## Requeriments
---
 * Go 1.19 or high
 * Node.js 16.9 or lower
 * Docker
## Run
---
* Backend
``` 
    go mod tidy
    go run main.go

    -- or 

    go build -o api main.go
    ./api
```
* Frontend
```
    npm install
    npm update
    npm start
```
* Database
```
    docker compose -f ./database.yaml up -d
```
## Routes
---
### Insert
 * Url: http://{IP}:8000/api/personalidades
 * Request Expected:
```
Method: POST
{
    "nome": "Name"
    "historia": "History"
} 
```
### Delete
 * Url: http://{IP}:8000/api/personalidades/{ip}
 * Request Expected:
```
Method: DELETE
{
    "nome": ""
    "historia": ""
} 
```
### Update
 * Url: http://{IP}:8000/api/personalidades/{ip}
 * Request Expected:
```
Method: PUT
{
    "nome": "New Name"
    "historia": "New History"
} 
```
### List
 * Url: http://{IP}:8000/api/personalidades/
 * Request Expected:
```
Method: GET
{
    "nome": ""
    "historia": ""
} 
```
### Get
 * Url: http://{IP}:8000/api/personalidades/{id}
 * Request Expected:
```
Method: GET
{
    "nome": ""
    "historia": ""
} 
```