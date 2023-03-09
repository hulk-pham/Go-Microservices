run:
	go build -o bin/main cmd/api/main.go && ./bin/main

dev:
	nodemon --exec go run cmd/api/main.go --signal SIGTERM

test:
	go test ./... | { grep -v 'no test files'; true; }

docker:
	docker-compose up -d --build
