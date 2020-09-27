package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(r http.Handler, method, path string, body io.Reader, headers map[string] string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}