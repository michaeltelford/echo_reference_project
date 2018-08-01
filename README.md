# Echo Reference Project

Starter project for developing RESTful API's using Go lang and the Echo HTTP web development framework.

Use this project as a reference or starter for your own applications.

## Branches

I've set up different repository branches to reflect what is provided for you out of the box. For example, the `master` branch contains only the Echo web router with no database connection. Other branches however, provide such functionality. So take a look and pick the one that suits you best for your needs.

This branch sees the API connect to and query a `postgres` database.

## Pre-requisites

- To use the database and API, you'll need to ensure `docker` and `docker-compose` are installed
- To use `make debug` you'll need to ensure `watcher` is in your `$PATH` (https://github.com/canthefason/go-watcher)
- To use `make convey` you'll need to ensure `goconvey` is in your `$PATH` (https://github.com/smartystreets/goconvey)
- To use `make lint` you'll need to ensure `golint` is in your `$PATH` (https://github.com/golang/lint)
- To use `make dep` you'll need to ensure `dep` in in your `$PATH` (`brew install dep`)

## Main Libraries

Type | Name | URL
---- | ---- | ----
Web | echo | https://echo.labstack.com/guide
Config | viper | https://github.com/spf13/viper
Database | sqlx | https://github.com/jmoiron/sqlx
Test Mocks | charlatan | https://github.com/percolate/charlatan

## Usage

Run `make help` for the full list of commands but in a nutshell:

```sh
$ make run
go run cmd/main.go

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
