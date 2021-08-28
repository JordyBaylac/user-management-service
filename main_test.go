package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/JordyBaylac/user-management-service/api"
	"github.com/JordyBaylac/user-management-service/api/handlers"
	"github.com/gofiber/fiber/v2"
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
			expectedBody: &handlers.CreateUserResponse{Email: "messi@gmail.com", Name: "Messi"},
			expectedCode: http.StatusOK,
		},
		{
			scenario:     "user email must be unique",
			requestBody:  handlers.CreateUserRequest{Email: "messi@gmail.com", Name: "Cristiano"},
			expectedCode: http.StatusConflict,
		},
	}

	// Setup the api
	app := api.Setup(nil)

	for _, test := range tests {
		// Arrange
		req := CreateUserHelper(app, test.requestBody)

		// Act
		res, _ := app.Test(req, -1)

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
		assert.Equalf(t, test.expectedBody.Name, response.Name, test.scenario, "Name must match")
		assert.NotEmpty(t, response.ID, test.scenario, "ID must have a value")
	}
}

func TestGetUser(t *testing.T) {

	validUserID := "valid-user-id"
	validUserName := "Modric"
	validUserEmail := "modric@real.madrid"

	tests := []struct {
		scenario string
		userID   string

		// Expected outputs
		expectedCode int
		expectedBody *handlers.GetUserResponse // optional
	}{
		{
			scenario:     "missing user id param should not be allowed",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			scenario:     "getting with not registered user id should not return anything",
			userID:       "invalid-user-id",
			expectedCode: http.StatusConflict,
		},
		{
			scenario:     "getting with registered user id should return user data",
			userID:       validUserID,
			expectedCode: http.StatusOK,
			expectedBody: &handlers.GetUserResponse{
				ID:    validUserID,
				Email: validUserEmail,
				Name:  validUserName,
			},
		},
	}

	// Setup the api
	generator := MockIDGenerator{validUserID}
	app := api.Setup(generator)
	createUserReq := CreateUserHelper(app, handlers.CreateUserRequest{Email: validUserEmail, Name: validUserName})
	app.Test(createUserReq, -1)

	for _, test := range tests {
		// Arrange
		httpRequest, _ := http.NewRequest(
			"GET",
			fmt.Sprintf("/users/%s", test.userID),
			nil,
		)
		httpRequest.Header.Add("Content-Type", "application/json")

		// Act (-1 disables request latency)
		res, _ := app.Test(httpRequest, -1)

		// Assert
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.scenario, "HTTP Status must match")

		if test.expectedBody == nil {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, test.scenario, "Body must not be empty")

		var response handlers.GetUserResponse
		err = json.Unmarshal(body, &response)
		assert.Nilf(t, err, test.scenario, "Body must be a valid JSON")

		assert.Equalf(t, test.expectedBody.ID, response.ID, test.scenario, "ID must match")
		assert.Equalf(t, test.expectedBody.Email, response.Email, test.scenario, "Email must match")
		assert.Equalf(t, test.expectedBody.Name, response.Name, test.scenario, "Name must match")
	}
}

func TestUpdateUser(t *testing.T) {

	validUserID := "valid-user-id"
	validUserName := "Modric"
	validUserEmail := "modric@real.madrid"
	newUserName := "Luka Modric"

	tests := []struct {
		scenario    string
		userID      string
		requestBody *handlers.UpdateUserRequest

		// Expected outputs
		expectedCode int
		expectedBody *handlers.UpdateUserResponse // optional
	}{
		{
			scenario:     "missing request body should not be allowed",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			scenario:     "missing user id param should not be allowed",
			requestBody:  &handlers.UpdateUserRequest{Name: newUserName},
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			scenario:     "missing required properties should not be allowed",
			userID:       validUserID,
			requestBody:  &handlers.UpdateUserRequest{Name: ""},
			expectedCode: http.StatusBadRequest,
		},
		{
			scenario:     "updating with not registered user id should not be allowed",
			userID:       "invalid-user-id",
			requestBody:  &handlers.UpdateUserRequest{Name: newUserName},
			expectedCode: http.StatusConflict,
		},
		{
			scenario:     "updating with registered user id should be ok",
			userID:       validUserID,
			requestBody:  &handlers.UpdateUserRequest{Name: newUserName},
			expectedCode: http.StatusOK,
			expectedBody: &handlers.UpdateUserResponse{
				ID:    validUserID,
				Email: validUserEmail,
				Name:  newUserName,
			},
		},
	}

	// Setup the api
	generator := MockIDGenerator{validUserID}
	app := api.Setup(generator)
	createUserReq := CreateUserHelper(app, handlers.CreateUserRequest{Email: validUserEmail, Name: validUserName})
	app.Test(createUserReq, -1)

	for _, test := range tests {
		// Arrange
		var requestBytes []byte
		if test.requestBody != nil {
			requestBytes, _ = json.Marshal(test.requestBody)
		}

		httpRequest, _ := http.NewRequest(
			"PATCH",
			fmt.Sprintf("/users/%s", test.userID),
			bytes.NewReader(requestBytes),
		)
		httpRequest.Header.Add("Content-Type", "application/json")

		// Act (-1 disables request latency)
		res, _ := app.Test(httpRequest, -1)

		// Assert
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.scenario, "HTTP Status must match")

		if test.expectedBody == nil {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		assert.Nilf(t, err, test.scenario, "Body must not be empty")

		var response handlers.UpdateUserResponse
		err = json.Unmarshal(body, &response)
		assert.Nilf(t, err, test.scenario, "Body must be a valid JSON")

		assert.Equalf(t, test.expectedBody.ID, "invalid", test.scenario, "ID must match")
		assert.Equalf(t, test.expectedBody.Email, response.Email, test.scenario, "Email must match")
		assert.Equalf(t, test.expectedBody.Name, response.Name, test.scenario, "Name must match")
	}
}

func CreateUserHelper(app *fiber.App, body handlers.CreateUserRequest) *http.Request {
	requestBytes, _ := json.Marshal(body)
	httpRequest, _ := http.NewRequest(
		"POST",
		"/users",
		bytes.NewReader(requestBytes),
	)
	httpRequest.Header.Add("Content-Type", "application/json")

	return httpRequest
}

type MockIDGenerator struct {
	userID string
}

func (generator MockIDGenerator) GenerateID() string {
	return generator.userID
}
