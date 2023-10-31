all: build
build:
	go build -o bin/app

run: build
	./bin/app

test:
	go test -v avelino/element -count=1
	