package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/models"
)

// GetAllSeries obtiene todas las series de la base de datos
func GetAllSeries() ([]models.Series, error) {
	query := `
		SELECT id, title, description, genre, year, rating, image_url, watched, created_at, updated_at
		FROM series
		ORDER BY created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying series: %w", err)
	}
	defer rows.Close()

	var seriesList []models.Series
	for rows.Next() {
		var s models.Series
		err := rows.Scan(
			&s.ID,
			&s.Title,
			&s.Description,
			&s.Genre,
			&s.Year,
			&s.Rating,
			&s.ImageURL,
			&s.Watched,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning series: %w", err)
		}
		seriesList = append(seriesList, s)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating series: %w", err)
	}

	return seriesList, nil
}

// GetSeriesByID obtiene una serie por su ID
func GetSeriesByID(id int) (*models.Series, error) {
	query := `
		SELECT id, title, description, genre, year, rating, image_url, watched, created_at, updated_at
		FROM series
		WHERE id = $1
	`

	var s models.Series
	err := DB.QueryRow(query, id).Scan(
		&s.ID,
		&s.Title,
		&s.Description,
		&s.Genre,
		&s.Year,
		&s.Rating,
		&s.ImageURL,
		&s.Watched,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Serie no encontrada
	}
	if err != nil {
		return nil, fmt.Errorf("error querying series by id: %w", err)
	}

	return &s, nil
}

// CreateSeries crea una nueva serie en la base de datos
func CreateSeries(req *models.CreateSeriesRequest) (*models.Series, error) {
	query := `
		INSERT INTO series (title, description, genre, year, rating, image_url, watched)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, title, description, genre, year, rating, image_url, watched, created_at, updated_at
	`

	var s models.Series
	err := DB.QueryRow(
		query,
		req.Title,
		req.Description,
		req.Genre,
		req.Year,
		req.Rating,
		req.ImageURL,
		req.Watched,
	).Scan(
		&s.ID,
		&s.Title,
		&s.Description,
		&s.Genre,
		&s.Year,
		&s.Rating,
		&s.ImageURL,
		&s.Watched,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating series: %w", err)
	}

	return &s, nil
}

// UpdateSeries actualiza una serie existente
func UpdateSeries(id int, req *models.UpdateSeriesRequest) (*models.Series, error) {
	query := `
		UPDATE series
		SET title = $1, description = $2, genre = $3, year = $4, 
		    rating = $5, image_url = $6, watched = $7, updated_at = $8
		WHERE id = $9
		RETURNING id, title, description, genre, year, rating, image_url, watched, created_at, updated_at
	`

	var s models.Series
	err := DB.QueryRow(
		query,
		req.Title,
		req.Description,
		req.Genre,
		req.Year,
		req.Rating,
		req.ImageURL,
		req.Watched,
		time.Now(),
		id,
	).Scan(
		&s.ID,
		&s.Title,
		&s.Description,
		&s.Genre,
		&s.Year,
		&s.Rating,
		&s.ImageURL,
		&s.Watched,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil // Serie no encontrada
	}
	if err != nil {
		return nil, fmt.Errorf("error updating series: %w", err)
	}

	return &s, nil
}

// DeleteSeries elimina una serie por su ID
func DeleteSeries(id int) error {
	query := `DELETE FROM series WHERE id = $1`

	result, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting series: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // Serie no encontrada
	}

	return nil
}
