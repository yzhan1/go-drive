package main

import (
	"fmt"
	"godrive/handler"
	"net/http"
)

func main() {
	fmt.Println("Server live!")

	http.HandleFunc("/files/upload", handler.UploadHandler)
	http.HandleFunc("/files/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/files/search", handler.FileQueryHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s", err.Error())
	}
}
