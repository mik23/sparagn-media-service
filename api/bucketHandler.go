package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
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

	ctx := appengine.NewContext(c.Request)

	_, err = service.GetBucketFactory(ctx, service.Google).Put(uploadedFile, bucketName, f)
	u, err := url.Parse("/" + bucketName + "/" + uploadedFile.Filename)
	if err != nil {
		util.ShowError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}

//Download saves the content
func Download(c *gin.Context) {
	fileName := c.Query("fileName")

	bucketName := "bucket-categ"
	object, err := service.GetBucketFactory(c, service.Google).Get(fileName, bucketName)

	var bytes []byte = nil
	if err == nil {
		bytes, _ = ioutil.ReadAll(object)
	}

	if err != nil {
		util.ShowError(c, err)
		return
	}

	c.Data(200, http.DetectContentType(bytes), bytes)
}
