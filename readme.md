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
- CQRS & Clean architecture
- gRPC
- Azure Blob
- Kafka pubsub
- Zinc (FTS)


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


## More
- Learn more: https://golangbot.com/learn-golang-series/
- Advance topic: https://www.golangprograms.com/advance-programs.html
- Log/Monitor/Healthcheck
- Background job
- Multiple thread
- Microservice
- Done Go Cloud Native