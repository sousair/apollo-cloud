package httputils

import (
	"mime/multipart"
	"os"
)

func MultipartHeaderFileToOsFile(headerFile *multipart.FileHeader) (*os.File, error) {
	file, err := headerFile.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	osFile, ok := file.(*os.File)

	if !ok {
		return nil, err
	}

	return osFile, nil
}
