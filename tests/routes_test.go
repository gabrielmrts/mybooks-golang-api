package integration

import (
	"net/http/httptest"
	"testing"

	"github.com/gabrielmrts/mybooks-golang-api/internal/api/routes"
	"github.com/stretchr/testify/assert"
)

func TestIntegration_Routes_Hello(t *testing.T) {
	router := routes.SetupRouter()

	req := httptest.NewRequest("GET", "/hello", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code)
}
