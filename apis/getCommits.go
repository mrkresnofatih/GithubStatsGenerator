package apis

import (
	json "encoding/json"
	"io"
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func GetCommits(repo string) (*[]GetCommitResponse, error) {
	token := os.Getenv("GH_TOKEN")
	login := os.Getenv("LOGIN")

	client := &http.Client{}
	url := "https://api.github.com/repos/" + login + "/" + repo + "/commits"
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleIfError(err)

	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)
	resp, err2 := client.Do(req)
	utils.HandleIfError(err2)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err3 := io.ReadAll(resp.Body)
		utils.HandleIfError(err3)

		data := []GetCommitResponse{}
		json.Unmarshal([]byte(body), &data)

		for _, value := range data {
			log.Println("Obtained Commit: " + value.Sha)
		}

		return &data, nil
	} else {
		e := utils.BuildError("getCommitsApi: StatusCode not 200 OK")
		return nil, e
	}
}
