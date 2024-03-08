package s3

import "github.com/sousair/apollo-cloud/internal/domain/repositories"

type S3FileRepository struct {
}

var _ repositories.FileRepository = (*S3FileRepository)(nil)

func NewS3FileRepository() *S3FileRepository {
	return &S3FileRepository{}
}

func (r *S3FileRepository) Upload(params repositories.UploadFileParams) (*repositories.Location, error) {
	location := &repositories.Location{
		URL: "https://s3.amazonaws.com/apollo-cloud/" + params.File.Name(),
	}
	return location, nil
}
