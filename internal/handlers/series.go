package handlers

import (
	"encoding/json"
	"net/http"
)

// SeriesHandler maneja GET /series y POST /series
func SeriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO: Implementar en Fase 2
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "GET /series - Coming soon in Phase 2",
		})
	case http.MethodPost:
		// TODO: Implementar en Fase 2
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "POST /series - Coming soon in Phase 2",
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SeriesDetailHandler maneja GET /series/:id, PUT /series/:id, DELETE /series/:id
func SeriesDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO: Implementar en Fase 2
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "GET /series/:id - Coming soon in Phase 2",
		})
	case http.MethodPut:
		// TODO: Implementar en Fase 2
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "PUT /series/:id - Coming soon in Phase 2",
		})
	case http.MethodDelete:
		// TODO: Implementar en Fase 2
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "DELETE /series/:id - Coming soon in Phase 2",
		})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
