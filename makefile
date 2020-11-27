DB_URI?=
FILE?=
NUM_EXECUTIONS?=

setup-env:
	docker-compose -f ./docker/docker-compose.yml up -d

build:
	go build -o main *.go

run:
	DB_URI=${DB_URI} go run *.go ${FILE}

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
