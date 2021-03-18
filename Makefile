.DEFAULT_GOAL := test

deps:
	go mod vendor

clean:
	rm -r bin

test:
	go test -v -timeout 30s -count=1 .