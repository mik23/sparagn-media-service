package service

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
	"sparagn.com/sparagn-media-service/util"
)

func CopyFile(c *gin.Context, sw *storage.Writer, file multipart.File) bool {
	if _, err := io.Copy(sw, file); err != nil {
		util.ShowError(c, err)
		return true
	}
	return false
}

func GetInstanceBucketClient(ctx context.Context, uploadedFile *multipart.FileHeader, bucketName string) (error, *storage.Writer) {
	path := util.RootDir()+"/resources/GCP/credentials/google.json"
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(path))

	writer := storageClient.Bucket(bucketName).Object(uploadedFile.Filename).NewWriter(ctx)
	return err, writer
}