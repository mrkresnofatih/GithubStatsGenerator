package apis

type GetGitTreesResponse struct {
	Sha  string     `json:"sha"`
	Tree []TreeNode `json:"tree"`
}

type TreeNode struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
}
