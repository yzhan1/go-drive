package db

import (
	"fmt"
	db "github.com/yzhan1/go-drive/db/mysql"
	"time"
)

type UserFile struct {
	Username   string
	FileHash   string
	FileSize   int64
	FileName   string
	UploadAt   string
	LastUpdate string
}

func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	statement, err := db.GetConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`)" +
			" values (?, ?, ?, ?, ?)")

	if err != nil {
		return false
	}
	defer statement.Close()

	_, err = statement.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func QueryUserFileMetadata(username string, limit int) ([]UserFile, error) {
	statement, err := db.GetConn().Prepare(
		"select file_sha1, file_name, file_size, upload_at, last_update from tbl_user_file" +
			" where user_name = ? limit ?")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(username, limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var userFiles []UserFile
	for rows.Next() {
		file := UserFile{}
		err = rows.Scan(&file.FileHash, &file.FileName, &file.FileSize, &file.UploadAt, &file.LastUpdate)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		userFiles = append(userFiles, file)
	}
	return userFiles, nil
}
