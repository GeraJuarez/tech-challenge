# tech-challenge

## Requirements:
* golang:1.16


# API Documentation

## URL

`http://localhost:8080/api`

## Hello

### `GET` Hello

The endpoint `GET /hello` returns `"Hello, World."`

## Pokemon

### `GET` Pokemon

The endpoint `GET /v1/pokemon/{name}` uses the external API `pokeapi.co ` to retrieve basic information about a single pokemon by name or pokedex ID


## How to run (from command line):
```bash
git clone <repo-name> # download source code from github
cd tech-challenge
cd src/
go mod download
go test ./... # run unit tests
go run . # run main program
```

## How to run (with docker):
```bash
git clone <repo-name> # download source code from github
cd tech-challenge
docker build --tag api src/ -f .
docker run --detach --publish 8080:8080 api
```
