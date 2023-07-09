it:
	go mod tidy

run:
	go run cli/unoserver.go

build: bin/unoserver

bin/unoserver:
	GOOS=linux go build -o build/unoserver-linux cli/unoserver.go
	GOOS=darwin go build -o build/unoserver-darwin cli/unoserver.go

clean:
	rm bin/* | true
