package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpload(t *testing.T) {
	t.Run("Should upload succeffully any file", func(t *testing.T) {
		path := "./uploader.go"
		file, err := os.Open(path)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		io.Copy(part, file)
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		res := httptest.NewRecorder()

		Upload(res, req)

		if res.Code != http.StatusOK {
			t.Error("not 200")
		}

		t.Log(res.Body.String())
	})

	t.Run("Should not upload if there is a missing file parameter", func(t *testing.T) {
		path := "./uploader.go"
		file, err := os.Open(path)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("randomKey", filepath.Base(path))

		io.Copy(part, file)
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		res := httptest.NewRecorder()

		Upload(res, req)

		if res.Code != http.StatusOK {
			t.Error("not 200")
		}

		isError := strings.Contains(res.Body.String(), "Error Retrieving the File")

		if isError == false {
			t.Error("Error not found")
		}

	})
}
