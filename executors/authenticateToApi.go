package executors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func Authenticate() {
	res, err := apis.GetUser()
	utils.HandleIfError(err)
	os.Setenv("LOGIN", res.Login)
}
