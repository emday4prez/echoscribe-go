build-EchoScribeFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	cp bootstrap $(ARTIFACTS_DIR)