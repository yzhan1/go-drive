package metadata

type FileMetadata struct {
	FileHash string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var metaDataMap map[string]FileMetadata

func init() {
	metaDataMap = make(map[string]FileMetadata)
}

func UpdateFileMetadata(data FileMetadata) {
	metaDataMap[data.FileHash] = data
}

func GetFileMetadata(fileHash string) FileMetadata {
	return metaDataMap[fileHash]
}
