package apis

import (
	json "encoding/json"
	io "io"
	log "log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	strconv "strconv"
)

func GetRepositories(token string, latest bool, page int, per_page int) (*[]GetRepositoriesResponse, error) {
	client := &http.Client{}
	query := "&sort=pushed&direction=desc"
	pageQuery := "&page=" + strconv.Itoa(page) + "&per_page=" + strconv.Itoa(per_page)
	if !latest {
		query = ""
	}
	url := "https://api.github.com/user/repos?" + query + pageQuery
	req, err := http.NewRequest("GET", url, nil)
	utils.HandleIfError(err)

	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)
	resp, err := client.Do(req)
	utils.HandleIfError(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		utils.HandleIfError(err)

		data := []GetRepositoriesResponse{}
		json.Unmarshal([]byte(body), &data)

		for _, value := range data {
			log.Println("Obtained Repo: " + value.Name)
		}

		return &data, nil

	} else {
		e := utils.BuildError("getRepositoriesApi: StatusCode not 200 OK")

		return nil, e
	}
}
