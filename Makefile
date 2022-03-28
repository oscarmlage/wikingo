run:
	go run main.go

rundebug:
	go run main.go -d

build:
	go build -o bin/wikingo

clean:
	rm -f bin/wikingo
