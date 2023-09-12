package usecase

import (
	"archive/zip"
	"bytes"
	"fmt"
	minio "libs/storage-client/golang/go-minio"
	"os"
)

type UnzipFileUseCase struct {
	contextEnv  string
	MinioClient minio.MinioClient
}

func NewUnzipFileUseCase(contextEnv string) *UnzipFileUseCase {
	return &UnzipFileUseCase{
		contextEnv: contextEnv,
		MinioClient: *minio.NewMinioClient(
			os.Getenv("MINIO_ENDPOINT"),
			os.Getenv("MINIO_ACCESS_KEY"),
			os.Getenv("MINIO_SECRET_KEY"),
		),
	}
}

func getBucketName(contextEnv string, source string) string {
	return fmt.Sprintf("raw-%s-source-%s", contextEnv, source)
}

func (ufu *UnzipFileUseCase) Execute(uri string, partition string, source string) ([]string, error) {
	fileBytes, err := ufu.MinioClient.DownloadFile(uri)
	if err != nil {
		return []string{}, err
	}

	bucketName := getBucketName(ufu.contextEnv, source)

	uris, err := ufu.unzipAndUpload(fileBytes, bucketName, partition)
	if err != nil {
		return []string{}, err
	}

	return uris, nil
}

func (ufu *UnzipFileUseCase) unzipAndUpload(zipData []byte, bucketName, partition string) ([]string, error) {
	zipDataReader := bytes.NewReader(zipData)
	zipReader, err := zip.NewReader(zipDataReader, int64(len(zipData)))
	if err != nil {
		return []string{}, err
	}

	uriUploadedFiles := make([]string, 0)

	for _, zipFile := range zipReader.File {
		fileReader, err := zipFile.Open()
		if err != nil {
			return []string{}, err
		}
		defer fileReader.Close()

		buffer := make([]byte, zipFile.UncompressedSize64)
		_, err = fileReader.Read(buffer)
		if err != nil {
			return []string{}, err
		}

		// Upload the file content to Minio
		uri, err := ufu.MinioClient.UploadFile(bucketName, zipFile.Name, partition, buffer)
		if err != nil {
			return []string{}, err
		}
		uriUploadedFiles = append(uriUploadedFiles, uri)
	}

	return uriUploadedFiles, nil
}
