package common

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Session() *session.Session {
	config := AppConfig()

	s3Config := &aws.Config{
		Region:      aws.String(config.AwsRegion),
		Credentials: credentials.NewStaticCredentials(config.AwsKey, config.AwsSecret, ""),
	}
	s3Session := session.New(s3Config)
	return s3Session
}

func UploadS3(file *multipart.FileHeader) (string, error) {
	config := AppConfig()
	fmt.Println(config)
	upFile, err := os.Open(file.Filename)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	defer upFile.Close()

	s3Session := GetS3Session()

	var fileSize int64 = file.Size
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	fileKey := strconv.Itoa(int(NowMinisecond())) + path.Ext(file.Filename)
	_, err = s3.New(s3Session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(config.S3BucketName),
		Key:                  aws.String(fileKey),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ServerSideEncryption: aws.String("AES256"),
	})

	fmt.Println(err)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", config.S3BucketName, config.AwsRegion, fileKey), nil

}
