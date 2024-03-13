package repositories

import (
	"os"

	"github.com/sousair/apollo-cloud/internal/domain/valueobjects"
)

type (
	FileRepository interface {
		Upload(params UploadFileParams) (*valueobjects.FileLocation, error)
	}

	UploadFileParams struct {
		File   *os.File
		Public bool
	}
)
