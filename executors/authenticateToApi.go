package executors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func Authenticate() {
	os.Setenv("GH_TOKEN", "")
	os.Setenv("REPO", "")

	res, err := apis.GetUser()
	utils.HandleIfError(err)
	os.Setenv("LOGIN", res.Login)
}
