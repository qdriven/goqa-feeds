# Generate Code By OpenAPI V3 Spec

## Install ogen

```shell
go get github.com/ogen-go/ogen/cmd/ogen
```

## create generate file

create generate file ```generate.go```
```go
package project

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target petstore --clean petstore.yml
```

## run script to generate codes

```shell
 go generate ./...
```

## Server Bootstrap

```shell
curl -X "POST" -H "Content-Type: application/json" --data "{\"name\":\"Cat\"}" http://localhost:8080/pet
```

## Build

```shell
go build bootstrap/petstore_main.go
```

## Run

```shell
./petstore_main
```