package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/database"
	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/models"
)

// ErrorResponse representa una respuesta de error
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SeriesHandler maneja GET /series y POST /series
func SeriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetAllSeries(w, r)
	case http.MethodPost:
		handleCreateSeries(w, r)
	default:
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SeriesDetailHandler maneja GET /series/:id, PUT /series/:id, DELETE /series/:id
func SeriesDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID de la URL
	id, err := extractIDFromPath(r.URL.Path)
	if err != nil {
		sendError(w, "Invalid series ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		handleGetSeriesByID(w, r, id)
	case http.MethodPut:
		handleUpdateSeries(w, r, id)
	case http.MethodDelete:
		handleDeleteSeries(w, r, id)
	default:
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetAllSeries maneja GET /series con filtros opcionales
func handleGetAllSeries(w http.ResponseWriter, r *http.Request) {
	// Obtener parámetros de query
	genre := r.URL.Query().Get("genre")
	search := r.URL.Query().Get("search")

	var series []models.Series
	var err error

	// Si hay filtros, usar la función con filtros
	if genre != "" || search != "" {
		series, err = database.GetSeriesWithFilters(genre, search)
	} else {
		// Sin filtros, obtener todas
		series, err = database.GetAllSeries()
	}

	if err != nil {
		sendError(w, "Failed to fetch series", http.StatusInternalServerError)
		return
	}

	// Si no hay series, devolver array vacío en lugar de null
	if series == nil {
		series = []models.Series{}
	}

	sendJSON(w, series, http.StatusOK)
}

// handleGetSeriesByID maneja GET /series/:id
func handleGetSeriesByID(w http.ResponseWriter, r *http.Request, id int) {
	series, err := database.GetSeriesByID(id)
	if err != nil {
		sendError(w, "Failed to fetch series", http.StatusInternalServerError)
		return
	}

	if series == nil {
		sendError(w, "Series not found", http.StatusNotFound)
		return
	}

	sendJSON(w, series, http.StatusOK)
}

// handleCreateSeries maneja POST /series
func handleCreateSeries(w http.ResponseWriter, r *http.Request) {
	var req models.CreateSeriesRequest

	// Decodificar JSON del body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validar datos
	if err := req.Validate(); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear serie en la base de datos
	series, err := database.CreateSeries(&req)
	if err != nil {
		sendError(w, "Failed to create series", http.StatusInternalServerError)
		return
	}

	sendJSON(w, series, http.StatusCreated)
}

// handleUpdateSeries maneja PUT /series/:id
func handleUpdateSeries(w http.ResponseWriter, r *http.Request, id int) {
	var req models.UpdateSeriesRequest

	// Decodificar JSON del body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validar datos (reutilizamos la validación de CreateSeriesRequest)
	createReq := models.CreateSeriesRequest{
		Title:       req.Title,
		Description: req.Description,
		Genre:       req.Genre,
		Year:        req.Year,
		Rating:      req.Rating,
		ImageURL:    req.ImageURL,
		Watched:     req.Watched,
	}

	if err := createReq.Validate(); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Actualizar serie en la base de datos
	series, err := database.UpdateSeries(id, &req)
	if err != nil {
		sendError(w, "Failed to update series", http.StatusInternalServerError)
		return
	}

	if series == nil {
		sendError(w, "Series not found", http.StatusNotFound)
		return
	}

	sendJSON(w, series, http.StatusOK)
}

// handleDeleteSeries maneja DELETE /series/:id
func handleDeleteSeries(w http.ResponseWriter, r *http.Request, id int) {
	err := database.DeleteSeries(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			sendError(w, "Series not found", http.StatusNotFound)
			return
		}
		sendError(w, "Failed to delete series", http.StatusInternalServerError)
		return
	}

	// 204 No Content - eliminación exitosa sin body
	w.WriteHeader(http.StatusNoContent)
}

// extractIDFromPath extrae el ID numérico de la URL /series/:id
func extractIDFromPath(path string) (int, error) {
	// path será algo como "/series/123" o "/series/123/"
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, &models.ValidationError{Message: "Invalid path"}
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, &models.ValidationError{Message: "Invalid ID format"}
	}

	return id, nil
}

// sendJSON envía una respuesta JSON
func sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// sendError envía una respuesta de error en JSON
func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
	})
}
