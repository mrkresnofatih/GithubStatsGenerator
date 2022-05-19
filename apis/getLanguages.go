package apis

import (
	json "encoding/json"
	"io"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func GetLanguages(repo string, lang string) (*int64, error) {
	login := os.Getenv("LOGIN")
	token := os.Getenv("GH_TOKEN")
	client := &http.Client{}
	url := "https://api.github.com/repos/" + login + "/" + repo + "/languages"

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

		data := map[string]int64{}
		json.Unmarshal([]byte(body), &data)

		languageSize := data[lang]
		return &languageSize, nil
	}

	e := utils.BuildError("getLanguagesApi: StatusCode not 200 OK")
	return nil, e
}
