package executors

import (
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func Authenticate() {
	os.Setenv("GH_TOKEN", "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")
	os.Setenv("REPO", "gh-svg-upload-target")

	res, err := apis.GetUser()
	utils.HandleIfError(err)
	os.Setenv("LOGIN", res.Login)
}
