it:
	go mod tidy

run:
	go run cmd/unoserver.go

build: bin/unoserver

bin/unoserver:
	GOOS=linux go build -o build/unoserver-linux cmd/unoserver.go
	GOOS=darwin go build -o build/unoserver-darwin cmd/unoserver.go

clean:
	rm bin/* | true
