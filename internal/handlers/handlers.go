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

	// Парсим форму, чтобы получить файл
	err := r.ParseMultipartForm(10 << 20) // 10 МБ
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "Conversion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем уникальное имя файла для сохранения результата
	filename := fmt.Sprintf("result_%s.txt", time.Now().UTC().Format("20060102_150405"))
	path := filepath.Join(".", filename)

	err = os.WriteFile(path, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Failed to save result: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
