bin/segment: *.go
	go build -o bin/segment

bin/test-segment: *.go
	go test -c
	mv segment.test bin/test-segment

run: bin/segment
	./bin/segment

test: bin/test-segment
ifdef TEST
	bin/test-segment -test.run="$(TEST)"
else
	bin/test-segment -test.v
endif
