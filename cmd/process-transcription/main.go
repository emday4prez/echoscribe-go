package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

var transcribeClient *transcribe.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}
	transcribeClient = transcribe.NewFromConfig(cfg)
}

func HandleRequest(ctx context.Context, s3Event events.S3Event) error {
	transcriptsBucket := os.Getenv("TRANSCRIPTS_BUCKET_NAME")

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key
		fileURI := fmt.Sprintf("s3://%s/%s", bucket, key)
		jobName := fmt.Sprintf("transcription-job-%d-%s", time.Now().Unix(), key)

		_, err := transcribeClient.StartTranscriptionJob(context.TODO(), &transcribe.StartTranscriptionJobInput{
			TranscriptionJobName: &jobName,
			LanguageCode:         types.LanguageCodeEnUs,
			Media: &types.Media{
				MediaFileUri: &fileURI,
			},
			OutputBucketName: &transcriptsBucket,
		})

		if err != nil {
			return fmt.Errorf("error starting transcription job: %w", err)
		}

		fmt.Printf("Successfully started transcription job %s for file %s\n", jobName, fileURI)
	}
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
