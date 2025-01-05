package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"currency-service/tests/fixtures"
	"currency-service/tests/helpers"

	"github.com/stretchr/testify/assert"
)

func TestGetRateByDate(t *testing.T) {
	date := "2025-01-05"

	router, db, err := helpers.SetupTestEnvironment()
	if err != nil {
		t.Fatalf("Failed to setup test environment: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %v", err)
	}
	path := filepath.Join(dir, "..", "fixtures", "rates.json")
	err = fixtures.LoadRatesFromFile(db, path)
	if err != nil {
		t.Fatalf("Failed to load fixture data: %v", err)
	}

	req, _ := http.NewRequest("GET", "/rates/"+date, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "USD")
	assert.Contains(t, w.Body.String(), "Euro")
	assert.Contains(t, w.Body.String(), "1.2345")
	assert.Contains(t, w.Body.String(), date)
}
