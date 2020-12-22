package service

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sparagn.com/sparagn-media-service/util"
)

type minioBucket struct {
	context context.Context
}

func (bucket *minioBucket) Get(objectName string, bucketName string) (io.Reader, error) {
	minioClient, err := GetMinioInstance()
	if err != nil {
		return nil, err
	}
	return minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
}

func (bucket *minioBucket) Put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (int64, error) {
	minioClient, err := GetMinioInstance()
	if err == nil {
		contentType, err := util.GetFileContentType(file)
		info, err := minioClient.PutObject(bucket.context, bucketName, uploadedFile.Filename, file, -1, minio.PutObjectOptions{ContentType: contentType})
		return info.Size, err
	}

	return 0, err

}

// GetMinioInstance builds a mino instance
func GetMinioInstance() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT") //"play.min.io"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := false

	// Initialize minio client object.
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
}
