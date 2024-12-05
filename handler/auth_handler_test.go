package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"project/domain"
	"project/service"
	"testing"
)

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := service.AuthServiceMock{}
	authHandler := NewAuthController(&mockService, zap.NewNop())

	tests := []struct {
		name               string
		requestBody        interface{}
		arg1MockSetup      bool
		arg2MockSetup      error
		expectedStatusCode int
		expectedBody       Response
	}{
		{
			name: "Success Login",
			requestBody: domain.User{
				Email:    "test@example.com",
				Password: "password123",
			},
			arg1MockSetup:      true,
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusOK,
			expectedBody: Response{
				Status:  true,
				Message: "user authenticated",
			},
		},
		{
			name:               "Invalid Request Body",
			requestBody:        "invalid-json",
			arg1MockSetup:      false,
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: Response{
				Status:  false,
				Message: "invalid request body",
			},
		},
		{
			name: "Missing Username or Password",
			requestBody: domain.User{
				Email:    "",
				Password: "",
			},
			arg1MockSetup:      false,
			arg2MockSetup:      nil,
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody: Response{
				Status:  false,
				Message: "authentication failed",
			},
		},
		{
			name: "Authentication Failed Username",
			requestBody: domain.User{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:      false,
			arg2MockSetup:      errors.New("invalid username and/or password"),
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody: Response{
				Status:  false,
				Message: "invalid username and/or password",
			},
		},
		{
			name: "Authentication Failed Password",
			requestBody: domain.User{
				Email:    "test@example.com",
				Password: "wrongpassword",
			},
			arg1MockSetup:      false,
			arg2MockSetup:      errors.New("invalid username and/or password"),
			expectedStatusCode: http.StatusUnauthorized,
			expectedBody: Response{
				Status:  false,
				Message: "invalid username and/or password",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err, "Failed to marshal request body")

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("GET", "/login", bytes.NewBuffer(requestBody))

			mockService.On("Login", tt.requestBody).Once().Return(tt.arg1MockSetup, tt.arg2MockSetup)

			authHandler.Login(c)

			res := w.Result()
			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)

			var responseBody Response
			err = json.NewDecoder(w.Body).Decode(&responseBody)
			assert.NoError(t, err, "Failed to decode response body")
			assert.Equal(t, tt.expectedBody.Status, responseBody.Status)
			assert.Equal(t, tt.expectedBody.Message, responseBody.Message)
		})
	}
}
