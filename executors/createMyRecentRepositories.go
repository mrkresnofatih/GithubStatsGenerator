package executors

import (
	"log"
	directors "mrkresnofatih/githubstatsgenerator/directors"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	exec "os/exec"
)

func CreateMyRecentRepositories() {
	log.Println("Starting Process: Create myRecentRepositories.svg")
	repos := directors.ObtainRecentRepositories(3)
	directors.MakeRecentReposScript(repos)

	cmd := exec.Command("bash", "makeRecentRepositoriesScript.txt")
	err := cmd.Run()
	utils.HandleIfError(err)
	log.Println("Successfully created myRecentRepositories.svg")
}
