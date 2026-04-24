package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	MaxUploadSize = 1 << 20 // 1 MB
	UploadPath    = "./uploads"
)

type UploadResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

// UploadImageHandler maneja el upload de imágenes
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limitar el tamaño del request
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)

	// Parse multipart form
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		sendError(w, "File too large. Maximum size is 1MB", http.StatusBadRequest)
		return
	}

	// Obtener el archivo
	file, header, err := r.FormFile("image")
	if err != nil {
		sendError(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validar tipo de archivo
	contentType := header.Header.Get("Content-Type")
	if !isValidImageType(contentType) {
		sendError(w, "Invalid file type. Only JPG, PNG, GIF, and WEBP are allowed", http.StatusBadRequest)
		return
	}

	// Crear directorio de uploads si no existe
	if err := os.MkdirAll(UploadPath, os.ModePerm); err != nil {
		sendError(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	// Generar nombre único
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filepath := filepath.Join(UploadPath, filename)

	// Crear archivo de destino
	dst, err := os.Create(filepath)
	if err != nil {
		sendError(w, "Error saving file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copiar el archivo
	_, err = io.Copy(dst, file)
	if err != nil {
		sendError(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Obtener la URL base del servidor
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	host := r.Host
	imageURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, host, filename)

	// Responder con la URL
	response := UploadResponse{
		URL:      imageURL,
		Filename: filename,
	}

	sendJSON(w, response, http.StatusCreated)
}

// isValidImageType verifica que el tipo de archivo sea válido
func isValidImageType(contentType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/gif",
		"image/webp",
	}

	contentType = strings.ToLower(contentType)
	for _, validType := range validTypes {
		if contentType == validType {
			return true
		}
	}
	return false
}
