package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
)

func ShowError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"error":   true,
	})
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}