package handlers

import (
	"encoding/json"
	"html/template"
	"io"
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

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error parsing form"})
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error retrieving file"})
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error reading file"})
		return
	}

	content := string(data)

	result, err := service.DetectAndConvert(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error processing content"})
		return
	}

	filename := "result_" + time.Now().UTC().Format("20060102_150405") + ".txt"
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error saving result"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message:  "File processed successfully",
		Filename: filename,
		Result:   result,
	})
}
