package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String())
}

func TestMainHandlerCorrectRequest(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestMainHandlerUnsupportedCity(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/cafe?count=3&city=samara", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "wrong city value")
}
