package api

import (
	"currency-service/models"
	"currency-service/tests/fixtures"
	"currency-service/tests/helpers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRates(t *testing.T) {
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

	req, _ := http.NewRequest("GET", "/rates", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NotEmpty(t, w.Body.String())

	var rates []models.Rate

	err = json.Unmarshal(w.Body.Bytes(), &rates)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	assert.NotEmpty(t, rates)

	for _, rate := range rates {
		assert.NotEmpty(t, rate.Code)
		assert.NotEmpty(t, rate.Name)
		assert.NotZero(t, rate.Nominal)
		assert.NotZero(t, rate.Rate)
		assert.NotEmpty(t, rate.Date)
	}
}
