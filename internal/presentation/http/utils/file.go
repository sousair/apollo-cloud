package httputils

import (
	"io"
	"mime/multipart"
	"os"
)

func MultipartHeaderFileToOsFile(headerFile *multipart.FileHeader) (*os.File, error) {
	file, err := headerFile.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newFile, err := os.Create(headerFile.Filename)

	if err != nil {
		return nil, err
	}
	_, err = io.Copy(newFile, file)

	if err != nil {
		return nil, err
	}

	return newFile, nil
}
