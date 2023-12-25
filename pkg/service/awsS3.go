package service

import (
	"bytes"
	"fmt"
	"mime/multipart"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

func CreateSession(cfg *config.S3Bucket) *session.Session {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
		},
	))
	return sess
}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadImageToS3(file *multipart.FileHeader, sess *session.Session) (string, error) {

	image, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer image.Close()

	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart"),
		Key:    aws.String("product images/" + fileName),
		Body:   image,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}

func UploadFilesToS3(file bytes.Buffer, sess *session.Session) (string, error) {

	// files, err := os.Open(file.String())
	// if err != nil {
	// 	return "", err
	// }
	// defer files.Close()
	fmt.Println("##", bytes.Buffer(file))
	fileName := uuid.New().String()

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart"),
		Key:    aws.String("files/" + fileName),
		Body:   bytes.NewReader(file.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return upload.Location, nil
}
