package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // запрос к сервису

	countStr := req.URL.Query().Get("count")
	count, _ := strconv.Atoi(countStr)

	//необходимые проверки

	if assert.LessOrEqual(t, count, totalCount) != true {
		for _, cafe := range cafeList["moscow"] {
			fmt.Println(cafe)
		}
	}

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
	req := httptest.NewRequest("GET", "/cafe?count=4&city=Ryazan", nil) // запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	city := req.URL.Query().Get("city")

	// необходимые проверки
	if assert.Equal(t, "moscow", city) == false {
		t.Errorf("StatusCode: %d. Wrong city value", responseRecorder.Code)
	}
}
