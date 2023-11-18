package service

import (
	"fmt"
	"mime/multipart"

	"github.com/Vajid-Hussain/Mobile-Mart-ecommerce/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("mobile-mart-image"),
		Key:    aws.String(file.Filename),
		Body:   image,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return upload.Location, nil
}

// func DownloadObject(url string, sess *session.Session, cfg *config.S3Bucket) error {

// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	defer file.Close()

// 	downloader := s3manager.NewDownloader(sess)
// 	numBytes, err := downloader.Download(file,
// 		&s3.GetObjectInput{
// 			Bucket: aws.String(bucket),
// 			Key:    aws.String(item),
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
// 	return nil
// }

// func FileUpload(s3Session *s3.S3, img *multipart.FileHeader) {
// 	file, err := os.Open(string(img))

// 	_, err := s3Session.PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String("mobile-mart-image"),
// 		Key:    aws.String("product/image.jpg"),
// 		Body:   img,
// 	})

// 	if err != nil {
// 		fmt.Println("Error uploading file to S3:", err)
// 		return
// 	}

// 	fmt.Println("File uploaded successfully.")
// }
