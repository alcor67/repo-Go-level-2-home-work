all:
#	@make bench
#	go test --bench=. --benchmem  > benchmark.txt
#	go test --bench=. --benchmem 
#test:
#	@make test
#	go test -v
#	go test ./... -v	
	go run ./main.go

test:
#	@make test
#	go test -v
	go test ./... -v
#	go test ./Example_confDub_test.go -v

linter:
#	@make linter
	go vet --all
	golangci-lint run

bench:
	go test --bench=. --benchmem --benchtime=10000000x > benchmark.txt
	go test --bench=. --benchmem --benchtime=10000000x


.PHONY: run
run:

#	godoc -http=:6060
# Команда запускает сервер документации на порту 6060
# Открыв в браузере адрес 
#	http://127.0.0.1:6060/pkg/github.com/alcor67/repo-Go-level-2-home-work/doc/
#	go run ./main.go
#	go run ./main.go -dir C:/Source/Golang/Go-level-2/lesson-8/Homework-8
#	go run ./main.go -dir testfolder
#	go run ./main.go -dir testfolderempty 
	go run ./main.go -dir folder 
#	go run ./main.go -dir ./folder -d -h
#	go run ./main.go -dir ./folder
#	go run ./main.go -dir ./folder/tmp1

doc:
	godoc -http=:6060