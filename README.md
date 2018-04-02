# Echo Reference Project

Starter project for developing RESTful API's using Go lang and the echo
HTTP routing library.

Use this project as a reference or starter for your own applications.

## Branches

I've set up different repository branches to reflect what is provided for you out of the box. For example, the `master` branch contains only the echo web router with no database connection. Other branches however, provide such functionality. So take a look and pick the one that suits you best for your needs.

## Pre-requisites

- To use `make debug` you'll need to ensure `watcher` is in your `$PATH` (https://github.com/canthefason/go-watcher)
- To use `make lint` you'll need to ensure `golint` is in your `$PATH` (https://github.com/golang/lint)

## Main Libraries

- echo (https://echo.labstack.com/guide)

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

Now hit the API with:

```sh
$ curl localhost:8000/v1/greet/rodney | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    27  100    27    0     0   4347      0 --:--:-- --:--:-- --:--:--  4500
{
  "message": "Hello, rodney"
}
```
