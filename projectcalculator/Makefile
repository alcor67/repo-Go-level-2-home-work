all:
#	@make test
	go test ./calculator -v


mytest:
#	@make test
	go test ./calculator -v
	go test ./... -v -race

linter:
#	@make linter
	go vet --all
	golangci-lint run

.PHONY: run
run:
#	в Makefile команда CGO_ENABLED=1 совместимость с языком C
	CGO_ENABLED=1 && go run -race ./main.go -op "div" -x 9 -y 3
	./main.exe -op "**" -x 5 -y 3