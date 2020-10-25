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

func Download() {
	// 	if err != nil {
	// 		d.errorf("readFile: unable to open file from bucket %q, file %q: %v", d.bucketName, fileName, err)
	// 		return
	// }
	// defer rc.Close()
	// slurp, err := ioutil.ReadAll(rc)
	// if err != nil {
	// 		d.errorf("readFile: unable to read data from bucket %q, file %q: %v", d.bucketName, fileName, err)
	// 		return
	// }

	// fmt.Fprintf(d.w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
	// if len(slurp) > 1024 {
	// 		fmt.Fprintf(d.w, "...%s\n", slurp[len(slurp)-1024:])
	// } else {
	// 		fmt.Fprintf(d.w, "%s\n", slurp)
	// }
}
