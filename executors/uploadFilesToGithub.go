package executors

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
)

func UploadFilesToGithub() {
	repo := os.Getenv("REPO")

	_, err := apis.GetTargetRepositoryExists()
	utils.HandleIfError(err)

	fileNames := []string{"myTopLanguages.svg", "myRecentRepositories.svg", "myRecentLanguages.svg", "myRecentContributions.svg"}
	for _, val := range fileNames {
		f, err := os.Open(val)
		utils.HandleIfError(err)
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)

		path := "generated/" + val
		_, er := apis.PostFile(repo, path, encoded)
		utils.HandleIfError(er)

		log.Println("Uploaded File: " + val + " to Repo: " + repo)
	}
}
