package service

import (
	"context"
	"io"
	"mime/multipart"
)

//IBucketFactory inteface for the factory
type IBucketFactory interface {
	Get(objectName string, bucketName string) (io.Reader, error)
	Put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (int64, error)
}

//BucketType declaration
type BucketType int

//Bucket Type enums
const (
	Google = iota
	Amazon
	Minio
)

//GetBucketFactory provides the correct bucket struct
func GetBucketFactory(c context.Context, bucketType BucketType) IBucketFactory {
	switch bucketType {
	case Google:
		return &googleBucket{context: c}
	case Minio:
		return &minioBucket{context: c}
	default:
		return &minioBucket{context: c}
	}
}
