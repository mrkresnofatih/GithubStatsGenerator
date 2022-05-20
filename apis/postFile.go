package apis

import (
	"bytes"
	json "encoding/json"
	"io"
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func PostFile(repo string, path string, file string) (bool, error) {
	login := os.Getenv("LOGIN")
	token := os.Getenv("GH_TOKEN")

	fileSha := ""
	latestCommit, e := GetNewestCommit(repo)
	if e == nil {
		shaOfTreeLatestCommit := latestCommit.Commit.Tree.Sha

		gitTree, errr := GetGitTrees(repo, shaOfTreeLatestCommit)
		utils.HandleIfError(errr)
		log.Println("Successful GetGitTree Request for Repo: " + repo)

		trees := gitTree.Tree
		for _, val := range trees {
			if val.Path == path {
				fileSha = val.Sha
				log.Println("Path: " + path + " exists => UpdateFile Operation")
			}
		}
	}

	client := &http.Client{}
	url := "https://api.github.com/repos/" + login + "/" + repo + "/contents/" + path
	pkg := PostFileRequest{"Upload File " + path, file, fileSha}
	jsn, err := json.Marshal(pkg)
	utils.HandleIfError(err)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsn))
	utils.HandleIfError(err)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)

	resp, err := client.Do(req)
	utils.HandleIfError(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
		return true, nil
	}

	body, ex := io.ReadAll(resp.Body)
	utils.HandleIfError(ex)
	response := map[string]string{}
	json.Unmarshal([]byte(body), &response)

	for _, v := range response {
		log.Println(v)
	}

	er := utils.BuildError("Status Code not 201 Created")
	return false, er
}
