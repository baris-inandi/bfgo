build:
	rm -rf ./bin
	mkdir ./bin
	echo "building binary for linux amd64"
	GOOS=linux GOARCH=amd64 go build -o ./bin/brainfuck-linux-amd64 ./compiler/brainfuck.go
	echo "building binary for linux arm"
	GOOS=linux GOARCH=arm go build -o ./bin/brainfuck-linux-arm ./compiler/brainfuck.go

clean:
	rm -rf ./bin

install:
	sudo go build -o /usr/bin/brainfuck .
