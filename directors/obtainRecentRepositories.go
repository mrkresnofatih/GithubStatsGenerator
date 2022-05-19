package directors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func ObtainRecentRepositories(limit int) []apis.GetRepositoriesResponse {
	token := os.Getenv("GH_TOKEN")

	repos := []apis.GetRepositoriesResponse{}
	result := []apis.GetRepositoriesResponse{}

	per_page := 30
	page := 1
	resp, err := apis.GetRepositories(token, true, page, per_page)
	utils.HandleIfError(err)

	repos = append(repos, *resp...)

	// Filter Repos with No Programming Languages
	for _, value := range repos {
		if value.Language != "" {
			result = append(result, value)
		}
	}

	// Select Top Limit Recent
	if len(result) > limit {
		result = result[:limit]
	}

	// Handler If No Repositories Available, Highly Unlikely
	if len(result) == 0 {
		e := utils.BuildError("0 Repos collected! Will not proceed!")
		utils.HandleIfError(e)
	}

	return result
}
