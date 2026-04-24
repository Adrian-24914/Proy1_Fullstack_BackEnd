package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/database"
	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/handlers"
	"github.com/Adrian-24914/Proy1_Fullstack_BackEnd/internal/middleware"
)

func main() {
	// Obtener puerto de variable de entorno (para Railway)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicializar base de datos
	if err := database.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.Close()

	// Configurar rutas
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthCheckHandler)

	// Series endpoints
	mux.HandleFunc("/series", handlers.SeriesHandler)
	mux.HandleFunc("/series/", handlers.SeriesDetailHandler)

	// Upload endpoint ← NUEVO
	mux.HandleFunc("/upload", handlers.UploadImageHandler)

	// Servir archivos estáticos (imágenes subidas) ← NUEVO
	fs := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	// Aplicar middleware CORS
	handler := middleware.CORS(mux)

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

// healthCheckHandler verifica el estado del servidor y la base de datos
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar conexión a la base de datos
	dbStatus := "disconnected"
	if err := database.DB.Ping(); err == nil {
		dbStatus = "connected"
	}

	response := map[string]string{
		"status":   "ok",
		"database": dbStatus,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
