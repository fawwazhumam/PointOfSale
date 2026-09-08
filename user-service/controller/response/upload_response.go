package response

type UploadPhotoResponse struct {
	URL      string `json:"url"`
	Path     string `json:"urlpath"`
	Filename string `json:"filename"`
}
