BINARY=engine
run:
	go run cmd/main.go

test: 
	go test -v -cover -covermode=atomic ./...