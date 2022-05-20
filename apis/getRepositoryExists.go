package apis

import (
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func GetTargetRepositoryExists() (bool, error) {
	token := os.Getenv("GH_TOKEN")
	login := os.Getenv("LOGIN")
	repo := os.Getenv("REPO")

	log.Println("Checking If Target Repository Exists")
	client := &http.Client{}
	url := "https://api.github.com/repos/" + login + "/" + repo
	req, err := http.NewRequest(http.MethodGet, url, nil)
	utils.HandleIfError(err)
	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)
	resp, err := client.Do(req)
	utils.HandleIfError(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	e := utils.BuildError("Target Repository Does Not Exist!")
	return false, e
}
