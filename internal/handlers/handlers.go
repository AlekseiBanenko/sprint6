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
	http.ServeFile(w, r, "index.html") // Отдаем файл с диска[web:24]
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим multipart-форму[web:14][web:15]
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы "file"
	f, fh, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file in form", http.StatusBadRequest)
		return
	}
	defer f.Close() // Закрываем файл

	// Читаем содержимое файла
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

	// Создаем локальный файл: timestamp + ext оригинала
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Записываем результат
	if _, err := outFile.WriteString(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем результат клиенту
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
