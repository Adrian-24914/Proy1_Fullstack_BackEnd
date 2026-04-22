package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// Initialize inicializa la conexión a la base de datos
func Initialize() error {
	var err error

	// Obtener DATABASE_URL de variable de entorno
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Fallback para desarrollo local
		databaseURL = "postgresql://postgres:postgres@localhost:5432/series_tracker?sslmode=disable"
	}

	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Verificar conexión
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("✅ Database connected successfully")

	// Ejecutar migraciones
	if err = runMigrations(); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}

// Close cierra la conexión a la base de datos
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// runMigrations ejecuta las migraciones de la base de datos
func runMigrations() error {
	query := `
	CREATE TABLE IF NOT EXISTS series (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		genre VARCHAR(100),
		year INTEGER,
		rating DECIMAL(3, 1) DEFAULT 0.0,
		image_url TEXT,
		watched BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX IF NOT EXISTS idx_series_title ON series(title);
	CREATE INDEX IF NOT EXISTS idx_series_genre ON series(genre);
	`

	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("✅ Migrations completed successfully")
	return nil
}
