build:
	go build -o main *.go

run:
	go run *.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
