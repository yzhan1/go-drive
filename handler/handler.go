package handler

import (
	"encoding/json"
	"fmt"
	"github.com/yzhan1/go-drive/metadata"
	"github.com/yzhan1/go-drive/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		view, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "Internal Server Error")
			return
		}
		io.WriteString(w, string(view))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get file data: %s", err.Error())
			return
		}
		defer file.Close()

		fileMetaData := metadata.FileMetadata{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		localFile, err := os.Create(fileMetaData.Location)
		if err != nil {
			fmt.Printf("Failed to create file: %s", err.Error())
			return
		}
		defer localFile.Close()

		fileMetaData.FileSize, err = io.Copy(localFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file: %s", err.Error())
			return
		}

		localFile.Seek(0, 0)
		fileMetaData.FileHash = util.FileSha1(localFile)
		metadata.UpdateFileMetadata(fileMetaData)

		http.Redirect(w, r, "/files/upload/success", http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Upload Success!")
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0]
	fileMetadata := metadata.GetFileMetadata(fileHash)

	data, err := json.Marshal(fileMetadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form.Get("filehash")
	fileMetadata := metadata.GetFileMetadata(fileHash)

	f, err := os.Open(fileMetadata.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Description", "attachment; filename=\""+fileMetadata.FileName+"\"")
	w.Write(data)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	op := r.Form.Get("op")
	fileHash := r.Form.Get("filehash")
	newName := r.Form.Get("filename")

	if op != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fileMetadata := metadata.GetFileMetadata(fileHash)
	fileMetadata.FileName = newName
	metadata.UpdateFileMetadata(fileMetadata)

	data, err := json.Marshal(fileMetadata)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form.Get("filehash")
	fileMetadata := metadata.GetFileMetadata(fileHash)
	defer os.Remove(fileMetadata.Location)

	metadata.DeleteFileMetadata(fileHash)

	w.WriteHeader(http.StatusOK)
}
