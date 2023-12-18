package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	tests := []struct {
		body string
		want string
	}{
		{body: "", want: "Hello world!"},
		{body: "Hello", want: "Hello world!"},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", "/", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(test.body))
		w := httptest.NewRecorder()

		handle(w, req)

		if got := w.Body.String(); got != test.want {
			t.Errorf("handle(%q) = %q, want %q", test.body, got, test.want)
		}
	}
}
