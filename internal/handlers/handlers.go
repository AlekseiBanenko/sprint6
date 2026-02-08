package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := os.ReadFile(file.(*os.File).Name())
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	content := string(data)

	// Определение и конвертация
	result, err := service.DetectAndConvert(content)
	if err != nil {
		http.Error(w, "Error processing content", http.StatusInternalServerError)
		return
	}

	// Создание уникального имени файла
	filename := "result_" + time.Now().UTC().Format("20060102_150405") + ".txt"
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		http.Error(w, "Error saving result", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Результат:\n%s\n\nФайл сохранен как %s", result, filename)
}
