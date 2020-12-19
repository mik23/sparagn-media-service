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

	ctx := appengine.NewContext(c.Request)

	bucketName := "image-categories"

	written, err := GetBucketFactory(service.Google).put(ctx, bucketName, uploadedFile, f)

	u, err := url.Parse("/" + bucketName + "/" + writer.Attrs().Name)
	if err != nil {
		util.ShowError(c, err)
		return
	}
f
	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}

//Download saves the content
func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	ctx := appengine.NewContext(c.Request)

	bucketName := "image-categories"
	object, error := GetBucketFactory(service.Google).get(fileName, bucketName)

	if err == nil {
		return ioutil.ReadAll(object)
	}
	defer object.Close()

	if err != nil {
		util.ShowError(c, err)
		return
	}

	c.Data(200, object.contentType, data)
}
