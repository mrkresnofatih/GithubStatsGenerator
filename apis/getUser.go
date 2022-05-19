package apis

import (
	json "encoding/json"
	"io"
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	http "net/http"
	"os"
)

func GetUser() (*GetUserResponse, error) {
	token := os.Getenv("GH_TOKEN")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	auth_header_value := "token " + token
	req.Header.Add("Authorization", auth_header_value)
	resp, err := client.Do(req)
	utils.HandleIfError(err)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		utils.HandleIfError(err)

		data := GetUserResponse{}
		json.Unmarshal([]byte(body), &data)
		log.Println("Authenticated as: " + data.Login)

		return &data, nil
	} else {
		e := utils.BuildError("getUserApi: StatusCode not 200 OK")
		return nil, e
	}
}
