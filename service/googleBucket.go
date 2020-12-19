package service

import (
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"sparagn.com/sparagn-media-service/util"
)

type googleBucket struct {
	context *gin.Context
}

func (bucket *googleBucket) getInstanceBucketClient() (*storage.Client, error) {
	path := util.RootDir() + "/resources/GCP/credentials/google.json"
	return storage.NewClient(bucket.context, option.WithCredentialsFile(path))
}

func (bucket *googleBucket) get(objectName string, bucketName string) (io.Reader, error) {
	return bucket.getInstanceReader(objectName, bucketName)
}

func (bucket *googleBucket) put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (written int64, err error) {
	writer, err := bucket.getInstanceWriter(uploadedFile, bucketName)

	if err == nil {
		return io.Copy(writer, file)
	}

	return 0, err
}

func (bucket *googleBucket) getInstanceWriter(fileHeader *multipart.FileHeader, bucketName string) (*storage.Writer, error) {
	storageClient, error := bucket.getInstanceBucketClient()

	if error == nil {
		return storageClient.Bucket(bucketName).Object(fileHeader.Filename).NewWriter(bucket.context), nil
	}

	return nil, error
}

func (bucket *googleBucket) getInstanceReader(fileName string, bucketName string) (*storage.Reader, error) {
	storageClient, error := bucket.getInstanceBucketClient()
	if error != nil {
		return nil, error
	}

	return storageClient.Bucket(bucketName).Object(fileName).NewReader(bucket.context)
}
