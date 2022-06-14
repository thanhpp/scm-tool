package fileutil

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileUtil interface {
	SaveFilesFromMultipart(dir, suffix string, files []*multipart.FileHeader) ([]string, error)
	DeleteAllFiles(paths []string)
}

func NewFileUtil() FileUtil {
	return &fileUtilImpl{}
}

type fileUtilImpl struct{}

func (fu fileUtilImpl) SaveFilesFromMultipart(dir, suffix string, files []*multipart.FileHeader) ([]string, error) {
	var (
		fileNames     = make([]string, 0, len(files))
		shouldCleanUp = false
	)
	defer func() {
		if shouldCleanUp {
			fu.DeleteAllFiles(fileNames)
		}
	}()

	for i := range files {
		f, err := files[i].Open()
		if err != nil {
			shouldCleanUp = true
			return nil, err
		}
		tmpFile, err := os.CreateTemp(
			dir,
			fmt.Sprintf("%s-%d-*.%s", suffix, time.Now().UnixMilli(), filepath.Ext(files[i].Filename)))
		if err != nil {
			shouldCleanUp = true
			return nil, err
		}

		if _, err := io.Copy(tmpFile, f); err != nil {
			shouldCleanUp = true
			return nil, err
		}

		fileNames = append(fileNames, tmpFile.Name())

		f.Close()
		tmpFile.Close()
	}

	return fileNames, nil
}

func (fileUtilImpl) DeleteAllFiles(paths []string) {
	for i := range paths {
		os.Remove(paths[i])
	}
}
