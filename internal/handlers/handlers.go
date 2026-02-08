package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

type Response struct {
	Message  string `json:"message"`
	Filename string `json:"filename,omitempty"`
	Result   string `json:"result,omitempty"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Message: "Method Not Allowed"})
		return
	}

	// Проверка Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType == "" || !containsMultipart(contentType) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Content-Type must be multipart/form-data"})
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		log.Println("Error parsing multipart form:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error parsing form"})
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Error retrieving file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error retrieving file"})
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error reading file"})
		return
	}

	content := string(data)

	// Обработка содержимого через сервис
	result, err := service.DetectAndConvert(content)
	if err != nil {
		log.Println("Error processing content:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error processing content"})
		return
	}

	// Проверка и парсинг JSON-ответа, если есть
	if len(result) > 0 && result[0] == '{' {
		var resp map[string]string
		if err := json.Unmarshal([]byte(result), &resp); err == nil {
			if val, ok := resp["result"]; ok {
				result = val
			}
		}
	}

	// Создание уникального имени файла
	filename := "result_" + time.Now().UTC().Format("20060102_150405") + ".txt"
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		log.Println("Error saving result file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error saving result"})
		return
	}

	// Отправка успешного ответа
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message:  "File processed successfully",
		Filename: filename,
		Result:   result,
	})
}

// Вспомогательная функция для проверки Content-Type
func containsMultipart(contentType string) bool {
	return len(contentType) >= 19 && (contentType[:19] == "multipart/form-data")
}
