test:
	go test ./...

test-c:
	go test -cover -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

check:
	go fmt ./...
	go vet ./...
	golint ./...
	# honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...
