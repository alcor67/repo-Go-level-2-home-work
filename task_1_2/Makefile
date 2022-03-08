all:
#	@make bench
#	go test --bench=. --benchmem  > benchmark.txt
#	go test --bench=. --benchmem 
test:
#	@make test
	go test -v
#	go test ./... -v	

linter:
#	@make linter
	go vet --all
	golangci-lint run

bench:
#	go test --bench=. --benchmem --benchtime=10000000x > benchmark.txt
#	go test --bench=. --benchmem --benchtime=10000000x

.PHONY: run
run:
	go run ./main.go