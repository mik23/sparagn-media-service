package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"sparagn.com/sparagn-media-service/util"
)

type minioBucket struct {
	context *gin.Context
}

func (bucket *minioBucket) get(objectName string, bucketName string) (io.Reader, error) {
	minioClient, err := GetMinioInstance()
	if err == nil {
		return nil, err
	}
	return minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
}

func (bucket *minioBucket) put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (int64, error) {
	minioClient, err := GetMinioInstance()
	if err == nil {
		contentType, err := util.GetFileContentType(file)
		info, err := minioClient.PutObject(bucket.context, bucketName, uploadedFile.Filename, file, uploadedFile.Size, minio.PutObjectOptions{ContentType: contentType})
		return info.Size, err
	}

	return 0, err

}

// GetMinioInstance builds a mino instance
func GetMinioInstance() (*minio.Client, error) {
	endpoint := "localhost:9000" //"play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := false

	// Initialize minio client object.
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
}
