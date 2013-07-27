BIN=$(GOPATH)/bin

$(BIN)/segment: *.go
	go build -o $(BIN)/segment

$(BIN)/test-segment: *.go
	go test -c
	mv segment.test $(BIN)/test-segment

test: $(BIN)/test-segment
ifdef TEST
	$(BIN)/test-segment -test.run="$(TEST)"
else
	$(BIN)/test-segment -test.v
endif

format:
	gofmt -w *.go

lint:
	go get github.com/golang/lint/golint
	$(BIN)/golint *.go
	go vet

.PHONY: test format lint
