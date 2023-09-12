package gominio

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
}

func NewMinioClient(minioEndpoint string, minioAccessKey string, minioSecretKey string) *MinioClient {
	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	return &MinioClient{Client: client}
}

func (m *MinioClient) getObject(bucketName string, fileName string) ([]byte, error) {
	object, err := m.Client.GetObject(
		context.Background(),
		bucketName,
		fileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}
	defer object.Close()

	var objContent []byte
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, object)
	if err != nil {
		return nil, err
	}
	objContent = buf.Bytes()

	return objContent, nil
}

// "http://minio:9000/raw-br-source-ceaf/20230911/raw.zip"
func (m *MinioClient) DownloadFile(uri string) ([]byte, error) {
	// Parse the URI to extract the bucket name and object key
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// Extract bucket name and object key from the URI path
	pathParts := strings.Split(strings.TrimPrefix(parsedURI.Path, "/"), "/")
	if len(pathParts) < 2 {
		return nil, errors.New("invalid URI path")
	}

	bucketName := pathParts[0]
	objectKey := strings.Join(pathParts[1:], "/")

	// Download the object's content
	objContent, err := m.getObject(bucketName, objectKey)
	if err != nil {
		return nil, err
	}

	return objContent, nil
}

func (m *MinioClient) UploadFile(bucketName string, fileName string, partition string, fileContent []byte) (string, error) {
	fileSize := int64(len(fileContent))
	// 128MB Because Spark its optimal size
	partSize := int64(1024 * 1024 * 128) // 128MB

	extension := strings.Split(fileName, ".")[1]
	fileNameClean := strings.Split(fileName, ".")[0]

	partPath := fmt.Sprintf("%s/%s", fileNameClean, partition)

	numParts := (fileSize + partSize - 1) / partSize

	var offset int64
	var partNumber int

	for partNumber = 1; offset < numParts; partNumber++ {
		if partSize > (fileSize - offset) {
			partSize = fileSize - offset
		}

		// Upload a part.
		reader := bytes.NewReader(fileContent[offset : offset+partSize])
		partName := fmt.Sprintf("%s/part-%d.%s", partPath, partNumber, extension)
		_, err := m.Client.PutObject(context.Background(), bucketName, partName, reader, partSize, minio.PutObjectOptions{})
		if err != nil {
			return "", err
		}
		offset += partSize
	}

	// Check if offset equals fileSize after the loop
	if offset != fileSize {
		return "", errors.New("offset doesn't match fileSize")
	}
	return partPath, nil
}
