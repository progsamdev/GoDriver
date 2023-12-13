package main

import (
	"GoDriver/internal/bucket"
	"GoDriver/internal/queue"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func main() {

	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		panic(err)
	}

	ch := make(chan queue.QueueDto)
	qc.Consume(ch)

	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "aprenda-golang-drive-raw",
		BucketUpload:   "aprenda-golang-drive-gzip",
	}

	b := bucket.New(bucket.AwsProvider, bcfg)

	for msg := range ch {

	}

}
