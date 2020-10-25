package api

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"sparagn.com/sparagn-media-service/util/test"
)

func TestUpload(t *testing.T) {

	t.Run("Ping ok", func(t *testing.T) {
		r := SetupRouter()
		w := test.PerformRequest(r, "GET", "/ping", nil, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Should upload and download successfully any file", func(t *testing.T) {

		//Upload
		fileName := "router.go"
		file, err := os.Open(fileName)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(fileName))
		if err != nil {
			_ = writer.Close()
			t.Error(err)
		}

		io.Copy(part, file)
		writer.Close()

		r := SetupRouter()
		headers := map[string]string{"Content-Type": writer.FormDataContentType()}
		res := test.PerformRequest(r, "POST", "/upload", body, headers)

		if res.Code != http.StatusOK {
			t.Error("not 200")
		}

		//Download
		res = test.PerformRequest(r, "GET", "/download?fileName="+fileName, body, headers)

		if res.Code != http.StatusOK {
			t.Error("Image not found")
		}

	})

	t.Run("Should not upload if there is a missing file parameter", func(t *testing.T) {
		path := "./router.go"
		file, err := os.Open(path)
		if err != nil {
			t.Error(err)
		}

		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("randomParamKey", filepath.Base(path))

		io.Copy(part, file)
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r := SetupRouter()

		headers := map[string]string{"Content-Type": writer.FormDataContentType()}
		res := test.PerformRequest(r, "POST", "/upload", body, headers)

		type Response struct {
			Error   bool   `json:"error"`
			Message string `json:"message"`
		}

		var responseErr Response
		json.Unmarshal(res.Body.Bytes(), &responseErr)

		assert.Equal(t, responseErr.Message, "http: no such file")
		assert.True(t, responseErr.Error)

	})

}
