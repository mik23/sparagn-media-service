package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"sparagn.com/sparagn-media-service/service"
	"sparagn.com/sparagn-media-service/util"
)

// Upload uploads file to bucket
func Upload(c *gin.Context) {
	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		util.ShowError(c, err)
		return
	}

	defer f.Close()

	fmt.Printf("Uploaded File: %+v\n", uploadedFile.Filename)
	fmt.Printf("File Size: %+v\n", uploadedFile.Size)
	fmt.Printf("MIME Header: %+v\n", uploadedFile.Header)

	bucketName := "bucket-categ"

	_, err = service.GetBucketFactory(c, service.Minio).Put(uploadedFile, bucketName, f)

	if err != nil {
		util.ShowError(c, err)
		return
	}

	u, _ := url.Parse("/" + bucketName + "/" + uploadedFile.Filename)

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}

//Download saves the content
func Download(c *gin.Context) {
	fileName := c.Query("fileName")

	bucketName := "bucket-categ"
	object, err := service.GetBucketFactory(c, service.Minio).Get(fileName, bucketName)

	var bytes []byte = nil
	if err == nil {
		bytes, err = ioutil.ReadAll(object)
		if err != nil {
			c.Data(404, "application/json", nil)
		}
	}

	if err != nil {
		util.ShowError(c, err)
		return
	}

	c.Data(200, http.DetectContentType(bytes), bytes)
}
