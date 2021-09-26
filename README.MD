

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







