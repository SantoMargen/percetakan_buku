package upload

type RequestPathUpload struct {
	Path string `json:"path"`
}
type RequestUpload struct {
	IDFile   string `json:"id_file"`
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
	Path     string `json:"path"`
}
