package file

type File struct {
	ID       int64  `json:"id"`
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
}

type FileInfo struct {
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
}
