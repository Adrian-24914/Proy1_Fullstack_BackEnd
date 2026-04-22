package main

import (
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
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Series endpoints (los implementaremos en Fase 2)
	mux.HandleFunc("/series", handlers.SeriesHandler)
	mux.HandleFunc("/series/", handlers.SeriesDetailHandler)

	// Aplicar middleware CORS
	handler := middleware.CORS(mux)

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
