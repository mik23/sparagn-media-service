package service

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"sparagn.com/sparagn-media-service/util"
)

func CopyFile(c *gin.Context, sw *storage.Writer, file multipart.File) bool {
	if _, err := io.Copy(sw, file); err != nil {
		util.ShowError(c, err)
		return true
	}
	return false
}

func GetInstanceBucketClient(ctx context.Context) (error, *storage.Client) {
	path := util.RootDir() + "/resources/GCP/credentials/google.json"
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(path))
}

func GetInstanceWriter(ctx context.Context, uploadedFile *multipart.FileHeader, bucketName string) (error, *storage.Writer) {
	storageClient, err := GetInstanceBucketClient(ctx)
	writer := storageClient.Bucket(bucketName).Object(uploadedFile.Filename).NewWriter(ctx)
	return err, writer
}

func GetInstanceReader(ctx context.Context, bucketName string, filename string) (error, *storage.Reader) {
	storageClient, err := GetInstanceBucketClient(ctx)
	reader := storageClient.Bucket(bucketName).Object(fileName).NewReader(ctx)
	return err, reader
}
