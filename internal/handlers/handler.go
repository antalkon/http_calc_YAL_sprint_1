package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/antalkon/http_calc_YAL_sprint_1/internal/models"
	"github.com/antalkon/http_calc_YAL_sprint_1/internal/services"
)

// Проверка на наличие ключа в JSON
func validateRequestBody(r *http.Request) (*models.RequestBody, error) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req models.RequestBody
	if err := decoder.Decode(&req); err != nil {
		return nil, errors.New("Invalid input format")
	}

	// Проверка на пустое значение
	if strings.TrimSpace(req.Expression) == "" {
		return nil, errors.New("Expression is not valid")
	}

	return &req, nil
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Method not allowed"})
		return
	}

	// Валидация JSON-запроса
	reqBody, err := validateRequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: err.Error()})
		return
	}

	// Вычисление результата
	result, err := services.Calc(reqBody.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Expression is not valid"})
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ResponseBody{Result: result})
}
