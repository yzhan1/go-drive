package main

import (
	"fmt"
	"github.com/yzhan1/go-drive/handler"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/files/upload", handler.UploadHandler)
	http.HandleFunc("/files/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/files/search", handler.QueryHandler)
	http.HandleFunc("/files/query", handler.QueryRecentFileHandler)
	http.HandleFunc("/files/download", handler.DownloadHandler)
	http.HandleFunc("/files/update", handler.UpdateHandler)
	http.HandleFunc("/files/delete", handler.DeleteHandler)

	http.HandleFunc("/users/signup", handler.SignUpHandler)
	http.HandleFunc("/users/signin", handler.SignInHandler)
	http.HandleFunc("/users/info", handler.UserInfoHandler)

	fmt.Println("Server live on port 8080!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s", err.Error())
	}
}
