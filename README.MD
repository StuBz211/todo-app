TODO APP API

Installing migrate tools
https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md


start postgres in docker
```sh
docker run --name=todo-db -e POSTGRES_PASSWORD=postgres -p 5436:5432 --rm postgres
```

migrations
```
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' up
```
&
```
migrate -path ./schema -database 'postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable' down
```


start App
```
go run cmd/main.go
```

SignUp - create new user
```
curl --location --request POST 'http://localhost:8000/auth/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "test", "password": "test_password", "name": "user"}'
```



SignIn - get auth token
```
curl --location --request POST 'http://localhost:8000/auth/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{"username": "test", "password": "test_password"}'
```


Create todolist

```
curl --location --request POST 'http://localhost:8000/api/lists' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{"Title": "todo list 1", "Description": "Second todo list "}'
```


Get todolist

```
curl --location --request GET 'http://localhost:8000/api/lists' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```

Get by Id
```
curl --location --request GET 'http://localhost:8000/api/lists/<list_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```

Update 
```
curl --location --request PUT 'http://localhost:8000/api/lists/<list_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{"Title": "New todo List"}'
```

Delete 
```
curl --location --request DELETE 'http://localhost:8000/api/lists/<list_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>''
```


Work with list items
See all items
```
curl --location --request GET 'http://localhost:8000/api/lists/<list_id>/items' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```

Create Items 
```
curl --location --request POST 'http://localhost:8000/api/lists/<list_id>/items' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{"Title": "Buy Something"}'
```
Get item
```
curl --location --request GET 'http://localhost:8000/api/items/<item_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```
Update item

```
curl --location --request PUT 'http://localhost:8000/api/items/<item_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data-raw '{"Title": "Buy Something", "Done": true}'
```

Delete item 
```
curl --location --request DELETE 'http://localhost:8000/api/items/<item_id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```
