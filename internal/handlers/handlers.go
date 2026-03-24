package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Page loading error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Page Display error", http.StatusInternalServerError)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Form parsing error", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error receiving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "File reading error", http.StatusInternalServerError)
		return
	}

	input := strings.TrimSpace(string(data))

	result, err := service.MorseOrText(input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Convert error: %v", err), http.StatusBadRequest)
		return
	}

	timestamp := time.Now().UTC().String()
	timestamp = strings.ReplaceAll(timestamp, ":", "-")
	timestamp = strings.ReplaceAll(timestamp, ".", "-")

	ext := filepath.Ext(handler.Filename)

	outputFilename := fmt.Sprintf("%s%s", timestamp, ext)

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(result); err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
