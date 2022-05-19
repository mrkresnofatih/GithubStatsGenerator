package apis

type GetRepositoriesResponse struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	PushedAt    string   `json:"pushed_at"`
	Size        int64    `json:"size"`
	Language    string   `json:"language"`
	Topics      []string `json:"topics"`
	Visibility  string   `json:"visibility"`
}
