clean:
	go clean -testcache
test: clean
	go test ./... -coverprofile=c.out

