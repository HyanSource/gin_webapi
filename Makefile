hello:
	echo "hello world"

build:
	go build main.go

run:
	go run main.go

test:
	go test

testginget:
	go test -v -run=TestGinGet

testginpost:
	go test -v -run=TestGinPost