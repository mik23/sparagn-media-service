package service

import (
	"context"
	"io"
	"io/ioutil"
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

func ReadFile(c *gin.Context, rc *storage.Reader) []byte {
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		util.ShowError(c, err)
	}
	return data
}

func GetInstanceBucketClient(ctx context.Context) (*storage.Client, error) {
	path := util.RootDir() + "/resources/GCP/credentials/google.json"
	return storage.NewClient(ctx, option.WithCredentialsFile(path))
}

func GetInstanceWriter(ctx context.Context, uploadedFile *multipart.FileHeader, bucketName string) (error, *storage.Writer) {
	storageClient, error := GetInstanceBucketClient(ctx)
	return error, storageClient.Bucket(bucketName).Object(uploadedFile.Filename).NewWriter(ctx)
}

func GetInstanceReader(ctx context.Context, bucketName string, fileName string) (*storage.Reader, error) {
	storageClient, error := GetInstanceBucketClient(ctx)
	if error != nil {
		return nil, error
	}

	return storageClient.Bucket(bucketName).Object(fileName).NewReader(ctx)
}
