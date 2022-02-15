.PHONY: darwin
darwin:
#	darwin компиляция MacOS
#	вывод данных о скомпилированном файле file m 
	GOOS=darwin go build
	file m
	ls -la

.PHONY: linux
linux:
#	darwin компиляция linux
#	вывод данных о скомпилированном файле file m 
	GOOS=linux go build
	file m
	ls -la

.PHONY: windows
#	darwin компиляция MS Windows
#	вывод данных о скомпилированном файле file m.exe
windows:
	GOOS=windows go build
	file m.exe
	ls -la