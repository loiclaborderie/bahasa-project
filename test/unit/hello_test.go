package unit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/loiclaborderie/go-gin-template/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Switch to test mode
	gin.SetMode(gin.TestMode)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create a Gin context
	_, r := gin.CreateTestContext(w)

	// Add the handler to be tested
	r.GET("/hello", handlers.HelloHandler)

	// Create a request
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", response["message"])
}
