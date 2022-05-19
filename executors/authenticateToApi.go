package executors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func Authenticate() {
	os.Getenv("GH_TOKEN")

	res, err := apis.GetUser()
	utils.HandleIfError(err)

	os.Setenv("LOGIN", res.Login)
}
