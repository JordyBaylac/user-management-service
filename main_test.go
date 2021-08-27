package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/JordyBaylac/user-management-service/api"
	"github.com/JordyBaylac/user-management-service/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		scenario    string
		requestBody handlers.CreateUserRequest

		// Expected outputs
		expectedCode int
		expectedBody *handlers.CreateUserResponse // optional
	}{
		{
			scenario:     "invalid email should not be allowed",
			requestBody:  handlers.CreateUserRequest{Email: "messi@"},
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario:     "missing required properties should not be allowed",
			requestBody:  handlers.CreateUserRequest{Name: "Messi"},
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario:     "valid request should return newly created user",
			requestBody:  handlers.CreateUserRequest{Email: "messi@gmail.com", Name: "Messi"},
			expectedCode: http.StatusOK,
		},
		{
			scenario:     "user email must be unique",
			requestBody:  handlers.CreateUserRequest{Email: "messi@gmail.com", Name: "Cristiano"},
			expectedCode: http.StatusConflict,
		},
	}

	// Setup the api
	app := api.Setup()

	for _, test := range tests {
		// Arrange
		requestBytes, err := json.Marshal(test.requestBody)
		httpRequest, _ := http.NewRequest(
			"POST",
			"/users",
			bytes.NewReader(requestBytes),
		)
		httpRequest.Header.Add("Content-Type", "application/json")

		// Act (-1 disables request latency)
		res, err := app.Test(httpRequest, -1)

		// Assert
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.scenario, "HTTP Status must match")

		if test.expectedBody == nil {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, test.scenario, "Body must not be empty")

		var response handlers.CreateUserResponse
		err = json.Unmarshal(body, &response)
		assert.Nilf(t, err, test.scenario, "Body must be a valid JSON")

		assert.Equalf(t, test.expectedBody.Email, response.Email, test.scenario, "Email must match")
		assert.Equalf(t, test.expectedBody.Name, response.Email, test.scenario, "Name must match")
		assert.NotEmpty(t, response.ID, test.scenario, "ID must have a value")
	}
}
