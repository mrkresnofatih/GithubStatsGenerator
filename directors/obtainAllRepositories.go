package directors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
	"time"
)

func ObtainAllRepositories() []apis.GetRepositoriesResponse {
	token := os.Getenv("GH_TOKEN")

	repos := []apis.GetRepositoriesResponse{}
	run := true
	page := 1
	perPage := 30
	for run {
		resp, err := apis.GetRepositories(token, true, page, perPage)
		utils.HandleIfError(err)

		repos = append(repos, *resp...)
		if len(*resp) == perPage {
			page++
		} else {
			run = false
		}

		time.Sleep(1 * time.Second)
	}

	return repos
}
