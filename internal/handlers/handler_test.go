package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antalkon/http_calc_YAL_sprint_1/internal/models"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name         string
		payload      string
		expectedCode int
		expectedBody interface{}
	}{
		{"Valid Expression", `{"expression":"3 + 5"}`, http.StatusOK, models.ResponseBody{Result: 8}},
		{"Division by Zero", `{"expression":"5 / 0"}`, http.StatusUnprocessableEntity, models.ErrorResponse{Error: "Expression is not valid"}},
		{"Invalid Character", `{"expression":"3 + a"}`, http.StatusUnprocessableEntity, models.ErrorResponse{Error: "Expression is not valid"}},
		{"Empty Expression", `{"expression":""}`, http.StatusUnprocessableEntity, models.ErrorResponse{Error: "Expression is not valid"}},
		{"Invalid JSON", `{"expr":"3 + 5"}`, http.StatusUnprocessableEntity, models.ErrorResponse{Error: "Invalid input format"}},
		{"Invalid Method", ``, http.StatusMethodNotAllowed, models.ErrorResponse{Error: "Method not allowed"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(test.payload))
			if test.name == "Invalid Method" {
				req = httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
			}
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			CalculateHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			// Проверка кода ответа
			if res.StatusCode != test.expectedCode {
				t.Errorf("Expected status %d, got %d", test.expectedCode, res.StatusCode)
			}

			// Проверка тела ответа
			var body interface{}
			if res.StatusCode == http.StatusOK {
				body = &models.ResponseBody{}
			} else {
				body = &models.ErrorResponse{}
			}
			if err := json.NewDecoder(res.Body).Decode(body); err != nil {
				t.Errorf("Failed to decode response body: %v", err)
			}

			expectedJSON, _ := json.Marshal(test.expectedBody)
			actualJSON, _ := json.Marshal(body)
			if string(expectedJSON) != string(actualJSON) {
				t.Errorf("Expected body: %s, got: %s", string(expectedJSON), string(actualJSON))
			}
		})
	}
}
