# Echo Reference Project

Starter project for developing RESTful API's using Golang and the Echo HTTP web development framework.

Use this project as a reference or starter for your own applications.

## Branches

I've set up different repository branches to reflect what is provided for you out of the box. For example, the `master` branch contains only the Echo web router with no database connection. Other branches however, provide such functionality. So take a look and pick the one that suits you best for your needs.

This branch sees the API connect to and query a `postgres` database.

## Main Golang Libraries

Type | Name | URL
---- | ---- | ----
Web | echo | https://echo.labstack.com/guide
Config | viper | https://github.com/spf13/viper
Database | sqlx | https://github.com/jmoiron/sqlx
DB Migrations | golang-migrate | https://github.com/golang-migrate/migrate
DB Mocks | go-sqlmock | https://github.com/DATA-DOG/go-sqlmock
Interface Mocks | charlatan | https://github.com/percolate/charlatan
Test Assertions | testify | https://github.com/stretchr/testify/assert

## Pre-requisites

### Required Packages

- `go` (with a correctly configured `$GOPATH` and `$GOPATH/bin` added to `$PATH`)
- `dep` (install with `brew install dep`)
- `docker`
- `docker-compose`

### Optional Packages

- To use `make convey` you'll need to `go get github.com/smartystreets/goconvey`
- To use `make lint` you'll need to `go get github.com/golang/lint`
- To use `make migrate` locally you'll need to install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cli)

### Required `ENV` Vars

Set the following vars in your environment:

```sh
VERSION=1
DEBUG=true
HOST=0.0.0.0
PORT=8000
DB_HOST=db
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=api
```

## Usage

Run `make help` for the full list of commands but in a nutshell:

```sh
$ make run
docker-compose up

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v3.3.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8000
```

Now create an author with:

```sh
$ curl -X POST \
>   http://localhost:8000/v1/authors \
>   -H 'Cache-Control: no-cache' \
>   -H 'Content-Type: application/json' \
>   -d '{
> "name": "John Smith",
> "age": 23,
> "salary": 25000
> }' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   102  100    50  100    52   7341   7634 --:--:-- --:--:-- --:--:--  8666
{
  "id": 1,
  "name": "John Smith",
  "age": 23
}
```

Now retrieve the author with:

```sh
$ curl -X GET \
>   http://localhost:8000/v1/authors/1 \
>   -H 'Cache-Control: no-cache' | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    50  100    50    0     0   7716      0 --:--:-- --:--:-- --:--:--  8333
{
  "id": 1,
  "name": "John Smith",
  "age": 23
}
```
