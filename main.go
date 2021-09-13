package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const LocalstackAddr = "localstack:4566"

func main() {
	waitForLocalStack()

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("eu-west-1"),
		Endpoint:         aws.String("http://" + LocalstackAddr),
		Credentials: credentials.NewStaticCredentials("test", "test", "test"),
		S3ForcePathStyle: aws.Bool(true),
	})

	if err != nil {
		log.Fatal(err)
	}

	svc := s3.New(sess)
	res, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d buckets\n", len(res.Buckets))
}

func waitForLocalStack() {
	start := time.Now()
	defer func() {
		fmt.Printf("It took %dms for localstack to be ready\n", time.Since(start).Milliseconds())
	}()

	fmt.Println("Checking if localstack is ready...")
	for i := 0; i < 15; i++ {
		fmt.Printf("Attempt %d \n", i)
		conn, err := net.Dial("tcp", LocalstackAddr)

		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			conn.Close()
			fmt.Println("Connected to localstack")
			break
		}
	}
}
