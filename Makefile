build-EchoScribeFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/hello
	cp bootstrap $(ARTIFACTS_DIR)

build-GenerateUploadUrlFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap-generate-url ./cmd/generate-upload-url
	cp bootstrap-generate-url $(ARTIFACTS_DIR)