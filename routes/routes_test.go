package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "hello, world!", rr.Body.String())
}

func TestFooHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FooHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "bar", rr.Body.String())
}

func TestRandomHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RandomHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	s, err := strconv.ParseFloat(rr.Body.String(), 64)
	if err != nil {
		t.Fatal(err)
	}
	assert.GreaterOrEqual(t, 1.0, s)
	assert.LessOrEqual(t, 0.0, s)
}

func TestEchoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "echo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EchoHandler)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code)
}
