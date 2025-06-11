package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/larryokubasu5460/gomart/user-service/handler"
	"github.com/larryokubasu5460/gomart/user-service/service"
)

func TestRegister(t *testing.T) {
	// simulate Gin Context
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := &service.UserService{} // methods can be mocked if needed
	h := &handler.UserHandler{Service: mockService}
	router.POST("/register", h.Register)

	body := map[string]string{
		"username":"testuser",
		"email":"test@example.com",
		"password":"test123",
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST","/register",bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type","application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}