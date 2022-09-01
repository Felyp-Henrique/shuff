all: build_linux

clean:
	rm -f shuff

build_linux:
	env GOOS=linux GOARCH=amd64 go build -o shuff main.go