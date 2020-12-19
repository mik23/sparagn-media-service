package service

import (
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

//IBucketFactory inteface for the factory
type IBucketFactory interface {
	get(objectName string, bucketName string) (io.Reader, error)
	put(uploadedFile *multipart.FileHeader, bucketName string, file multipart.File) (int64, error)
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
func GetBucketFactory(bucketType BucketType, c *gin.Context) IBucketFactory {
	switch bucketType {
	case Google:
		return &googleBucket{context: c}
	case Minio:
		return &minioBucket{context: c}
	default:
		return &minioBucket{context: c}
	}
}
