package repositories

import "os"

type (
	FileRepository interface {
		Upload(params UploadFileParams) (*Location, error)
	}

	Location struct {
		URL      string
		Provider string
	}

	UploadFileParams struct {
		File   *os.File
		Public bool
	}
)
