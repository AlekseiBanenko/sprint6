package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// ✅ ПРОВЕРЯЕМ multipart
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") {
		http.Error(w, "multipart/form-data required", http.StatusBadRequest)
		return
	}

	// ✅ Парсим форму
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, fmt.Sprintf("parse error: %v", err), http.StatusInternalServerError)
		return
	}

	// ✅ Ищем файл в FormFile
	f, fh, err := r.FormFile("myFile")
	if err != nil {
		// DEBUG: показываем что есть в форме
		for name := range r.MultipartForm.File {
			fmt.Printf("Found file field: %s\n", name)
		}
		http.Error(w, fmt.Sprintf("no file 'file': %v", err), http.StatusBadRequest)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём файл
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()
	outFile.WriteString(result)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
