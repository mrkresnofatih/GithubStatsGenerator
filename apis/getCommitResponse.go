package apis

type GetCommitResponse struct {
	Sha    string            `json:"sha"`
	Commit CommitInformation `json:"commit"`
}

type CommitInformation struct {
	Message string       `json:"message"`
	Author  CommitAuthor `json:"author"`
	Tree    Tree         `json:"tree"`
}

type CommitAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type Tree struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}
