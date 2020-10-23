package client


import (

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
)
const (
	AWS_S3_REGION = "us-west-2"
	AWS_S3_BUCKET = "iaheroku-models"
)

var sess = connectAWS()

func connectAWS() *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION),
		},
	)
	if err != nil {
		panic(err)
	}
	return sess
}

type S3Client struct {

}
func (client *S3Client) UploadFile(file io.Reader, fileName string) error {
	// Upload
	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key: aws.String(fileName),
		Body: file,
	})
	if err != nil {
		log.Fatal("An error occurred uploading the file", err)
		return err
	} else {
		log.Println("File uploaded")
	}
	return nil
}