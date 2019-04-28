package db

import (
	"database/sql"
	"fmt"
	db "github.com/yzhan1/go-drive/db/mysql"
)

func UpdateFileMetadata(filehash string, filename string, filesize int64, fileaddr string) bool {
	statement, err := db.GetConn().Prepare("insert ignore into tbl_file (`file_sha1`, `file_name`, `file_size`," +
		"`file_addr`, `status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Error preparing statement: " + err.Error())
		return false
	}
	defer statement.Close()

	ret, err := statement.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println("Error preparing statement: " + err.Error())
		return false
	}
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash %s has been uploaded", filehash)
		}
		return true
	}
	return false
}

type File struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func GetFileMetadata(filehash string) (*File, error) {
	statement, err := db.GetConn().Prepare(
		"select file_sha1,file_addr,file_name,file_size from tbl_file " +
			"where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println("Error preparing statement: " + err.Error())
		return &File{}, err
	}
	defer statement.Close()

	file := File{}
	err = statement.QueryRow(filehash).Scan(&file.FileHash, &file.FileAddr, &file.FileName, &file.FileSize)
	if err != nil {
		fmt.Println("Error querying statement: " + err.Error())
		return &file, err
	}
	return &file, nil
}
