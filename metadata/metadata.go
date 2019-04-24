package metadata

import (
	"github.com/yzhan1/go-drive/db"
	"sort"
)

type FileMetadata struct {
	FileHash string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var metadataMap map[string]FileMetadata

func init() {
	metadataMap = make(map[string]FileMetadata)
}

func UpdateFileMetadata(data FileMetadata) {
	metadataMap[data.FileHash] = data
}

func UpdateFileMetadataDB(data FileMetadata) bool {
	return db.OnFileUploadFinished(data.FileHash, data.FileName, data.FileSize, data.Location)
}

func GetFileMetadataDB(fileHash string) (FileMetadata, error) {
	file, err := db.GetFileMeta(fileHash)
	if err != nil {
		return FileMetadata{}, err
	}
	fileMetadata := FileMetadata{
		FileHash: file.FileHash,
		FileName: file.FileName.String,
		FileSize: file.FileSize.Int64,
		Location: file.FileAddr.String,
	}
	return fileMetadata, nil
}

func GetFileMetadata(fileHash string) FileMetadata {
	return metadataMap[fileHash]
}

func DeleteFileMetadata(fileHash string) {
	delete(metadataMap, fileHash)
}

func GetRecentFileMetadata(count int) []FileMetadata {
	arr := make([]FileMetadata, len(metadataMap))
	for _, v := range metadataMap {
		arr = append(arr, v)
	}
	sort.Sort(SortByUploadTime(arr))
	return arr[0: count]
}
