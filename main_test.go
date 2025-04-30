package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", ping)
}

func TestPing(t *testing.T) {
	type testCase struct {
		method               string
		path                 string
		expectedResponseCode int
		expectedResponseBody string
	}

	cases := []testCase{
		{"GET", "/ping", http.StatusOK, "{\"message\":\"Welcome to CoffeeShop!\"}"},
	}

	for _, c := range cases {
		req := getRequest(c.method, c.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		if c.expectedResponseCode != w.Code {
			t.Errorf("%s: expected response code %d, got %d", c.method, c.expectedResponseCode, w.Code)
		}

		if c.expectedResponseBody != string(body) {
			t.Errorf("%s: expected response body %q, got %q", c.method, c.expectedResponseBody, string(body))
		}
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, "/ping", nil)
	return req
}
