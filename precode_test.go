package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=9&city=moscow", nil) // запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки

	res := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(res))

}
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	require.Equal(t, 200, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)

}
func TestMainHandlerWhenCityNotIsMoscow(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=UnExistsCity", nil) // запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())

}
