run:
	go build -o bin/main cmd/api/main.go && ./bin/main

dev:
	nodemon --exec go run cmd/api/main.go --signal SIGTERM
	
swagdoc:
	swag init --parseDependency --parseInternal

test:
	go test ./... | { grep -v 'no test files'; true; }

docker:
	docker-compose up -d --build

gqlgen:
	go run github.com/99designs/gqlgen generate

rpcgen:
	cd rpc && protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto
