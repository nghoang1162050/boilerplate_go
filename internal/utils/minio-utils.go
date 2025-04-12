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

func (f *MinioClient) List(ctx context.Context, prefix string) <-chan minio.ObjectInfo {
	objectCh := f.minioClient.ListObjects(ctx, f.bucketName, minio.ListObjectsOptions{Prefix: prefix, Recursive: true})
	return objectCh
}

func (f *MinioClient) Download(ctx context.Context, objectName string) (io.Reader, error) {
	opts := minio.GetObjectOptions{}
	object, err := f.minioClient.GetObject(ctx, f.bucketName, objectName, opts)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (f *MinioClient) Delete(ctx context.Context, objectName string) error {
	opts := minio.RemoveObjectOptions{}
	err := f.minioClient.RemoveObject(ctx, f.bucketName, objectName, opts)
	if err != nil {
		return err
	}
	return nil
}
