package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"sprint6/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html><head><title>Morse</title></head><body>
<h1>Morse Converter</h1><form action="/upload" method="post" enctype="multipart/form-data">
<input type="file" name="file" accept=".txt" required><button>Convert</button></form></body></html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	result, err := service.AutoConvert(string(data))
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fname := fmt.Sprintf("result_%s.txt", time.Now().UTC().Format("20060102_150405"))
	os.WriteFile(fname, []byte(result), 0644)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
