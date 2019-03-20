package metadata

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

func GetFileMetadata(fileHash string) FileMetadata {
	return metadataMap[fileHash]
}

func DeleteFileMetadata(fileHash string) {
	delete(metadataMap, fileHash)
}
