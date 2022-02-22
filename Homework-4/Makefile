all:
#	@make test
	go test ./workerpool -v


mytest:
#	@make test
	go test ./workerpool -v
	go test ./... -v

linter:
#	@make linter
	go vet --all
	golangci-lint run

.PHONY: run
run:
	go run ./main.go
