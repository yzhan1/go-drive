package metadata

import "time"

const baseFormat = "2006-01-02 15:04:05"

type SortByUploadTime []FileMetadata

func (a SortByUploadTime) Len() int {
	return len(a)
}

func (a SortByUploadTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a SortByUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(baseFormat, a[i].UploadAt)
	jTime, _ := time.Parse(baseFormat, a[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
