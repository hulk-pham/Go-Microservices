.PHONY: run dev swag_doc test docker gql_gen rpc_gen open_swagger

run:
	go build -o bin/main cmd/api/main.go && ./bin/main

dev:
	nodemon --exec go run cmd/api/main.go --signal SIGTERM 

devrpc:
	nodemon --exec go run cmd/grpc/server.go --signal SIGTERM
	
swag_doc:
	swag init -g cmd/api/main.go --parseDependency --parseInternal && cp -r docs presentation/http && rm -r docs

test:
	go test ./... | { grep -v 'no test files'; true; }

docker:
	docker-compose up -d --build

gql_gen:
	go run github.com/99designs/gqlgen generate

rpc_gen:
	cd ./presentation/rpc && rm -rf ./pb && mkdir ./pb  && protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

open_swagger:
	open http://localhost:8080/swagger/index.html