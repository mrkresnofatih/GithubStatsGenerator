package apis

type PostFileRequest struct {
	Message string `json:"message"`
	Content string `json:"content"`
	Sha     string `json:"sha"`
}
