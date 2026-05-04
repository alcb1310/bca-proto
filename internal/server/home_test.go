package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alcb1310/bca-proto/internal/server"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	s := server.Server{
		Port: 8080,
		DB:   nil,
	}

	server := httptest.NewServer(s.Router())
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	s.Router().ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "hello", rec.Body.String())
}
