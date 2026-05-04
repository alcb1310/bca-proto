clean:
	rm -rf bin

build: clean
	go build -o bin/bca-proto cmd/api/main.go

run: build
	./bin/bca-proto
