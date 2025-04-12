package utils

import (
	"context"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	minioClient *minio.Client
	bucketName  string
}

var FileClient *MinioClient

func (f *MinioClient) NewMinioClient() error {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		
		Secure: useSSL,
	})
	
	if err != nil {
		return err
	}

	f.minioClient = minioClient
	f.bucketName = os.Getenv("MINIO_BUCKET")
	return nil
}

func (f *MinioClient) Upload(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
	}
	
	_, err := f.minioClient.PutObject(ctx, f.bucketName, objectName, reader, objectSize, opts)
    if err != nil {
        return err
    }

    return nil
}
