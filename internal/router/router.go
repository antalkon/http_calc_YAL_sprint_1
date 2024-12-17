package router

import (
	"net/http"

	"github.com/antalkon/http_calc_YAL_sprint_1/internal/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handlers.CalculateHandler)
	return mux
}
