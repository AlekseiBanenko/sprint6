package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	contentType := r.Header.Get("Content-Type")
	if contentType == "" || !strings.HasPrefix(contentType, "multipart/") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Content-Type must be multipart/form-data"})
		return
	}

	err := r.ParseMultipartForm(10 << 20)
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

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error reading file"})
		return
	}

	content := string(data)

	// Вызов сервиса обработки
	result, err := service.DetectAndConvert(content)
	if err != nil {
		log.Println("Error processing content:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error processing content"})
		return
	}

	// Если результат — JSON-объект, извлечь поле "result"
	if len(result) > 0 && result[0] == '{' {
		var resp map[string]string
		if err := json.Unmarshal([]byte(result), &resp); err == nil {
			if val, ok := resp["result"]; ok {
				result = val
			}
		}
	}

	// Создаем имя файла
	filename := "result_" + time.Now().UTC().Format("20060102_150405") + ".txt"
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		log.Println("Error saving result file:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error saving result"})
		return
	}

	// В ответе возвращаем именно исходный текст, а не JSON
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}

// Вспомогательная функция для проверки Content-Type
func containsMultipart(contentType string) bool {
	return len(contentType) >= 19 && (contentType[:19] == "multipart/form-data")
}
