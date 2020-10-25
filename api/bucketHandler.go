package api

import (
	"fmt"
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
	err, writer := service.GetInstanceWriter(ctx, uploadedFile, bucketName)
	if err != nil {
		util.ShowError(c, err)
	}

	if service.CopyFile(c, writer, f) {
		return
	}

	if err := writer.Close(); err != nil {
		util.ShowError(c, err)
		return
	}

	u, err := url.Parse("/" + bucketName + "/" + writer.Attrs().Name)
	if err != nil {
		util.ShowError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	})
}

func Download(c *gin.Context) {
	fileName := c.Query("fileName")
	ctx := appengine.NewContext(c.Request)

	bucketName := "image-categories"
	reader, err := service.GetInstanceReader(ctx, bucketName, fileName)

	defer reader.Close()

	if err != nil {
		util.ShowError(c, err)
		return
	}

	data := service.ReadFile(c, reader)
	c.Data(200, reader.Attrs.ContentType, data)
}
