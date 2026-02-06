package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// ✅ ВАРИАНТ 1: Сначала парсим форму, потом ищем файл
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, fmt.Sprintf("parse form error: %v", err), http.StatusInternalServerError)
		return
	}

	f, fh, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("no file: %v", err), http.StatusBadRequest)
		return
	}
	defer f.Close()

	// Читаем файл
	data, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Конвертируем
	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаём локальный файл
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	outFile.WriteString(result)

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
