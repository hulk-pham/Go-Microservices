# Golang project

## Feature
- Gin
- JWT
- GORM
- Viper
- Swagger
- Upload S3
- Unit test
- GraphQL
- Redis
- CORS config
- Realtime websocket


## Start

1. Create a .env file at the root of the project and add the following:
   ```
   cp .env.example .env
   ```
2. Run `go mod download`.
4. Run `make run` or `make dev`.

## To run RPC

- Install Protocol Buffer Compiler https://grpc.io/docs/protoc-installation/
- Mac user only
```
brew install protoc-gen-go

brew install protoc-gen-go-grpc
```


## TODO
- learn more: https://golangbot.com/learn-golang-series/
- advance topic: https://www.golangprograms.com/advance-programs.html
- Rate limit
- Design pattern
- gRPC
- Realtime chat
- Log/Monitor/Healthcheck
- Add more library
- Microservice
- Done Go Cloud Native