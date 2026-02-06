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
	http.ServeFile(w, r, "index.html") // ✅ Файл с диска!
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. ✅ ПАРСИМ MULTIPART ФОРМУ
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 2. ✅ ПОЛУЧАЕМ ФАЙЛ "file"
	f, fh, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "no file in form", http.StatusBadRequest)
		return
	}
	defer f.Close()

	// 3. ✅ ЧИТАЕМ СОДЕРЖИМОЕ
	data, err := io.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. ✅ КОНВЕРТИРУЕМ
	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. ✅ СОЗДАЁМ ЛОКАЛЬНЫЙ ФАЙЛ
	ext := filepath.Ext(fh.Filename)
	filename := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)
	outFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// 6. ✅ ЗАПИСЫВАЕМ РЕЗУЛЬТАТ
	if _, err := outFile.WriteString(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 7. ✅ ОТДАЁМ КЛИЕНТУ
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
