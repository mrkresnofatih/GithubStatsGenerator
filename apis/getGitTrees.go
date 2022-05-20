package apis

import (
	json "encoding/json"
	"io"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func GetGitTrees(repo string, sha string) (*GetGitTreesResponse, error) {
	login := os.Getenv("LOGIN")
	token := os.Getenv("GH_TOKEN")

	client := &http.Client{}
	url := "https://api.github.com/repos/" + login + "/" + repo + "/git/trees/" + sha + "?recursive=0"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	utils.HandleIfError(err)
	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)

	resp, er := client.Do(req)
	utils.HandleIfError(er)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		utils.HandleIfError(err)

		data := GetGitTreesResponse{}
		json.Unmarshal([]byte(body), &data)

		return &data, nil
	}

	e := utils.BuildError("getGitTrees: StatusCode not 200 OK")
	return nil, e
}
