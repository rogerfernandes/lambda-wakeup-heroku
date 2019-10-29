package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWakeUpHeroku(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()
	request := Request{HerokuURLs: []string{testServer.URL}}

	response, err := WakeUpHeroku(request)

	assert.Equal(t, "\n"+testServer.URL+" - Respond: body", response.Log)
	assert.NoError(t, err)
}

func TestWakeUpHerokuFail(t *testing.T) {
	url := "haaa"
	request := Request{HerokuURLs: []string{url}}

	response, err := WakeUpHeroku(request)

	assert.Equal(t, "\n"+url+" - Failed!", response.Log)
	assert.NoError(t, err)
}

func BenchmarkWakeUpHeroku(b *testing.B) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()
	request := Request{HerokuURLs: []string{testServer.URL}}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		WakeUpHeroku(request)
	}
}
