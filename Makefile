build-EchoScribeFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/hello/main.go
	cp bootstrap $(ARTIFACTS_DIR)

build-GenerateUploadUrlFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap-generate-url ./cmd/generate-upload-url/main.go
	cp bootstrap-generate-url $(ARTIFACTS_DIR)

build-TranscriptionProcessorFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap-process-transcription ./cmd/process-transcription/main.go
	cp bootstrap-process-transcription $(ARTIFACTS_DIR)