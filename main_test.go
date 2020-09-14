package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloSnykconHandlerReturnsCorrectMessage(context *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		context.Fatal(err)
	}
	expectedBody := `Hello Snykcon 2020!!!`
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloSnykcon)
	handler.ServeHTTP(recorder, request)
	if recorder.Body.String() != expectedBody {
		context.Errorf("bad response body, wanted %v got %v",
			recorder.Body, expectedBody)
	}
}
