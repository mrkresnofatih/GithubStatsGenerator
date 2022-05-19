package apis

type GetUserResponse struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Location          string `json:"location"`
	Bio               string `json:"bio"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	PublicRepos       int    `json:"public_repos"`
	TotalPrivateRepos int    `json:"total_private_repos"`
}
