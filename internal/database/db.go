package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// Initialize inicializa la conexión a la base de datos
func Initialize() error {
	var err error

	// Cargar .env (solo en desarrollo local)
	godotenv.Load()

	// Obtener DATABASE_URL de variable de entorno
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL not set. Please create a .env file")
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

	// Insertar datos de ejemplo si la tabla está vacía
	var count int
	DB.QueryRow("SELECT COUNT(*) FROM series").Scan(&count)

	if count == 0 {
		seedQuery := `
		INSERT INTO series (title, description, genre, year, rating, image_url, watched) VALUES
		('Breaking Bad', 'A high school chemistry teacher turned methamphetamine producer.', 'Crime', 2008, 9.5, 'https://images.unsplash.com/photo-1574267432644-f610a13241d1?w=400', true)
		`
		DB.Exec(seedQuery)
		log.Println("✅ Sample data inserted")
	}

	return nil
}
