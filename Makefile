test:
	go test -v ./...

a.out:
	go build -v -o $@ .
.PHONY: a.out
