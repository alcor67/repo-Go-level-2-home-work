all:

# while true; do go run main.go; sleep 1; done
#	GOMAXPROCS=2 go run main.go 2>trace.out
#	go tool trace trace.out

#	go run -race main.go

	GOMAXPROCS=2 go run -race main.go 2>trace.out
	go tool trace trace.out


	