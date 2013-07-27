bin/segment: *.go
	go build -o bin/segment

run: bin/segment
	./bin/segment
