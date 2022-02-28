all:
#	@make test
	


mytest:
#	@make test
	
	go test ./... -v

linter:
#	@make linter
	go vet --all
	golangci-lint run

.PHONY: run
run:
	go run ./main.go
