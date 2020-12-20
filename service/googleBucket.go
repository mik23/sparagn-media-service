package service

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"sparagn.com/sparagn-media-service/util"
)

type googleBucket struct {
	context context.Context
}

func (bucket *googleBucket) getInstanceBucketClient() (*storage.Client, error) {
	path := util.RootDir() + "/resources/GCP/credentials/google.json"
	return storage.NewClient(bucket.context, option.WithCredentialsFile(path))
}

func (bucket *googleBucket) Get(objectName string, bucketName string) (io.Reader, error) {
	return bucket.getInstanceReader(objectName, bucketName)
}

func (bucket *googleBucket) Put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (int64, error) {
	writer, err := bucket.getInstanceWriter(uploadedFile, bucketName)
	defer writer.Close()

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
