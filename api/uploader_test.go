package api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sparagn.com/sparagn-media-service/util/test"
	"testing"
)


func TestUpload(t *testing.T) {

	t.Run("Ping ok", func (t *testing.T){
		r := SetupRouter()
		w := test.PerformRequest(r, "GET", "/ping", nil, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Should upload successfully any file", func(t *testing.T) {
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
			_ = writer.Close()
			t.Error(err)
		}

		io.Copy(part, file)
		writer.Close()

		r := SetupRouter()
		headers := map[string] string {"Content-Type": writer.FormDataContentType()}
		res := test.PerformRequest(r, "POST", "/upload", body, headers)

		if res.Code != http.StatusOK {
			t.Error("not 200")
		}
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
		part, err := writer.CreateFormFile("randomParamKey", filepath.Base(path))

		io.Copy(part, file)
		writer.Close()

		req := httptest.NewRequest("POST", "/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r := SetupRouter()

		headers := map[string] string {"Content-Type": writer.FormDataContentType()}
		res := test.PerformRequest(r, "POST", "/upload", body, headers)

		type Response struct {
			Error bool `json:"error"`
			Message string `json:"message"`
		}

		var responseErr Response
		json.Unmarshal(res.Body.Bytes(), &responseErr)

		assert.Equal(t, responseErr.Message, "http: no such file")
		assert.True(t, responseErr.Error)

	})
}
