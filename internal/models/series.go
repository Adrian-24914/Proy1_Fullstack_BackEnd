package models

import (
	"time"
)

// Series representa una serie de TV
type Series struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Year        int       `json:"year"`
	Rating      float64   `json:"rating"`
	ImageURL    string    `json:"image_url"`
	Watched     bool      `json:"watched"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateSeriesRequest representa los datos para crear una serie
type CreateSeriesRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	ImageURL    string  `json:"image_url"`
	Watched     bool    `json:"watched"`
}

// UpdateSeriesRequest representa los datos para actualizar una serie
type UpdateSeriesRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Genre       string  `json:"genre"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	ImageURL    string  `json:"image_url"`
	Watched     bool    `json:"watched"`
}

// Validate valida los datos de una serie
func (s *CreateSeriesRequest) Validate() error {
	if s.Title == "" {
		return &ValidationError{Field: "title", Message: "title is required"}
	}
	if s.Year < 1900 || s.Year > 2100 {
		return &ValidationError{Field: "year", Message: "year must be between 1900 and 2100"}
	}
	if s.Rating < 0 || s.Rating > 10 {
		return &ValidationError{Field: "rating", Message: "rating must be between 0 and 10"}
	}
	return nil
}

// ValidationError representa un error de validación
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
