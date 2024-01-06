package main

import (
	"GoDriver/internal/bucket"
	"GoDriver/internal/queue"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
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

	b, err := bucket.New(bucket.AwsProvider, bcfg)

	if err != nil {
		panic(err)
	}

	for msg := range ch {
		src := fmt.Sprintf("%s/%s", msg.Path, msg.Filename)
		dst := fmt.Sprintf("%d_%s", msg.ID, msg.Filename)
		file, err := b.Download(src, dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
			continue
		}

		body, err := io.ReadAll(file)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, err = zw.Write((body))
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		if err := zw.Close(); err != nil {
			log.Printf("ERROR %v", err)
			continue
		}

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		err = b.Upload(zr, src)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

		os.Remove(dst)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}

	}

}
