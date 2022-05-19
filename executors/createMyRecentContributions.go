package executors

import (
	"log"
	directors "mrkresnofatih/githubstatsgenerator/directors"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	exec "os/exec"
)

func CreateMyRecentContributions() {
	log.Println("Starting Process: Create myRecentContributions.svg")
	recentRepos := directors.ObtainCommitsThisWeek()
	directors.MakeRecentContributionsScript(recentRepos)

	cmd := exec.Command("bash", "makeRecentContributionsScript.txt")
	err := cmd.Run()
	utils.HandleIfError(err)
	log.Println("Successfully created myRecentContributions.svg")
}
