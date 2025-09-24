build-GenerateUploadUrlFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap-generate-url ./cmd/generate-upload-url
	cp bootstrap-generate-url $(ARTIFACTS_DIR)

build-TranscriptionProcessorFunction:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bootstrap-process-transcription ./cmd/process-transcription
	cp bootstrap-process-transcription $(ARTIFACTS_DIR)