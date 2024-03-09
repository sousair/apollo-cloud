package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
)

type S3FileRepository struct {
	uploader          *s3manager.Uploader
	privateBucketName string
	publicBucketName  string
}

var _ repositories.FileRepository = (*S3FileRepository)(nil)

func NewS3FileRepository(awsSession *session.Session, privateBucketName, publicBucketName string) *S3FileRepository {
	uploader := s3manager.NewUploader(awsSession)

	return &S3FileRepository{
		uploader,
		privateBucketName,
		publicBucketName,
	}
}

func (r *S3FileRepository) Upload(params repositories.UploadFileParams) (*repositories.Location, error) {
	ACL := "private"
	BucketName := r.privateBucketName

	if params.Public {
		ACL = "public-read"
		BucketName = r.publicBucketName
	}

	filename := params.File.Name()

	uploadInput := &s3manager.UploadInput{
		Bucket: &BucketName,
		Key:    &filename,
		Body:   params.File,
		ACL:    &ACL,
	}

	res, err := r.uploader.Upload(uploadInput)

	if err != nil {
		return nil, err
	}

	location := &repositories.Location{
		URL:      res.Location,
		Provider: "s3",
	}

	return location, nil
}
