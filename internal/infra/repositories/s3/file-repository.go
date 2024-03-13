package s3

import (
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sousair/apollo-cloud/internal/domain/repositories"
	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
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

func (r S3FileRepository) Upload(params repositories.UploadFileParams) (*valueobjects.FileLocation, error) {
	ACL := "private"
	BucketName := r.privateBucketName

	if params.Public {
		ACL = "public-read"
		BucketName = r.publicBucketName
	}

	filename := params.File.Name()
	extension := filepath.Ext(filename)

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

	location := &valueobjects.FileLocation{
		URL:       res.Location,
		Provider:  "s3",
		Extension: extension,
	}

	return location, nil
}
