package main

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// POST http://127.0.0.1:3000/
func TestHomePOSTEndpoint_WhenValidMethod_ReturnString(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	req := httptest.NewRequest(fiber.MethodPost, "/", nil)
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "POST", string(body))
}

// POST http://127.0.0.1:3000/
func TestHomePOSTEndpoint_WhenInvalidMethod_ReturnError(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	// Test using PATCH method instead of POST
	req := httptest.NewRequest(fiber.MethodPatch, "/", nil)
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode) // Expect error

	// Check response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "Method Not Allowed", string(body))
}

// GET http://127.0.0.1:3000/hello [Hello World.]
func TestHelloGETEndpoint_WhenNoParam_ReturnDefaultString(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	req := httptest.NewRequest(fiber.MethodGet, "/hello", nil)
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "Hello World.", string(body))
}

// GET http://127.0.0.1:3000/hello/there [Hello there.]
func TestHelloGETEndpoint_WhenHavePathParam_ReturnDefaultStringWithPathParam(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	req := httptest.NewRequest(fiber.MethodGet, "/hello/there", nil)
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "Hello there.", string(body))
}

// GET http://127.0.0.1:3000/hello/there?msg=!!! [Hello there!!!]
func TestHelloGETEndpoint_WhenHaveQueryAndPathParam_ReturnDefaultStringWithQueryAndPathParam(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	req := httptest.NewRequest(fiber.MethodGet, "/hello/there?msg=!!!", nil)
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "Hello there!!!", string(body))
}

// POST http://127.0.0.1:3000/jsonResponse
func TestJsonResponseEndpoint_WhenValidJson_ReturnJson(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	// Prepare JSON payload
	payload := fiber.Map{
		"name": "John",
	}
	payloadBytes, _ := json.Marshal(payload)

	// Create a POST request with JSON body
	req := httptest.NewRequest(fiber.MethodPost, "/jsonResponse", bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check json response
	var responseData fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	assert.NoError(t, err)
	assert.Equal(t, payload, responseData)
}

// POST http://127.0.0.1:3000/jsonResponse
func TestJsonResponseEndpoint_WhenInvalidJson_ReturnErrorJsonResponse(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	// Simulate an invalid JSON payload (trailing comma)
	invalidJson := `{"name": "John",}` // Note the trailing comma
	req := httptest.NewRequest(fiber.MethodPost, "/jsonResponse", bytes.NewReader([]byte(invalidJson)))
	req.Header.Set("Content-Type", "application/json")

	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Check json response
	var responseData fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	assert.NoError(t, err)
	assert.Equal(t, "Cannot parse JSON", responseData["error"])
}

// POST http://127.0.0.1:3000/uploadFile
func TestUploadFileEndpoint_WhenValidFile_ReturnFileName(t *testing.T) {
	// Initialize necessary setup before test [Use a test suite to automate]
	fiberApp := fiber.New()
	setupControllers(fiberApp)

	// Create a buffer to write our multipart form
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Create a file to upload
	file, createFileErr := os.Create("testfile.txt")
	assert.NoError(t, createFileErr)

	defer file.Close()
	defer os.Remove(file.Name()) // Clean up the file afterwards

	// Add the file to the multipart form
	_, createFormErr := writer.CreateFormFile("file", filepath.Base(file.Name()))
	assert.NoError(t, createFormErr)

	writer.Close()

	// Create a new HTTP request with the multipart form
	req := httptest.NewRequest(fiber.MethodPost, "/uploadFile", &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform the request
	resp, err := fiberApp.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check response body
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)

	assert.Equal(t, "Uploaded fileName: "+file.Name(), string(body))
}
