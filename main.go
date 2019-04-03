package main

import (
	"fmt"
	"github.com/yzhan1/go-drive/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/files/upload", handler.UploadHandler)
	http.HandleFunc("/files/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/files/search", handler.QueryHandler)
	http.HandleFunc("/files/query", handler.QueryRecentFileHandler)
	http.HandleFunc("/files/download", handler.DownloadHandler)
	http.HandleFunc("/files/update", handler.UpdateHandler)
	http.HandleFunc("/files/delete", handler.DeleteHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s", err.Error())
	} else {
		fmt.Println("Server live on port 8080!")
	}
}
