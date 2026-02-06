package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const logFile = "network.log"

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Server started at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		data, err := os.ReadFile(logFile)
		if err != nil {
			w.Write([]byte("Log file is empty\n"))
			return
		}
		w.Write(data)

	case http.MethodPost:
		body, _ := io.ReadAll(r.Body)

		entry := fmt.Sprintf("[%s] %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			string(body),
		)

		file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		file.WriteString(entry)
		file.Close()

		w.Write([]byte("Log entry added\n"))

	case http.MethodDelete:
		os.Remove(logFile)
		w.Write([]byte("Log file cleared\n"))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
