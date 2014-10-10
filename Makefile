test:
	go test -v .

test-integration:
	go test -v ./...

get-deps:
	go get github.com/codegangsta/martini
