package upload

type ResponseUpload struct {
	IDFile   string `json:"id_file"`
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
}
