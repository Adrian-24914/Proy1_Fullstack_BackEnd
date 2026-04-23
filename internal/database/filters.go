package database

import (
	"fmt"

	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/models"
)

// GetSeriesWithFilters obtiene series con filtros opcionales
func GetSeriesWithFilters(genre, search string) ([]models.Series, error) {
	query := `
		SELECT id, title, description, genre, year, rating, image_url, watched, created_at, updated_at
		FROM series
		WHERE 1=1
	`
	args := []interface{}{}
	argCount := 1

	// Filtro por género
	if genre != "" {
		query += fmt.Sprintf(" AND LOWER(genre) = LOWER($%d)", argCount)
		args = append(args, genre)
		argCount++
	}

	// Búsqueda por título o descripción
	if search != "" {
		query += fmt.Sprintf(" AND (LOWER(title) LIKE LOWER($%d) OR LOWER(description) LIKE LOWER($%d))", argCount, argCount)
		args = append(args, "%"+search+"%")
		argCount++
	}

	query += " ORDER BY created_at DESC"

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying series with filters: %w", err)
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
