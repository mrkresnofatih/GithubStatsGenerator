package executors

import (
	"log"
	directors "mrkresnofatih/githubstatsgenerator/directors"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	exec "os/exec"
)

func CreateMyRecentLanguages() {
	log.Println("Starting Process: Create myRecentLanguages.svg")
	langs := directors.ObtainRecentLanguages()
	directors.MakeRecentLanguagesScript(langs)

	cmd := exec.Command("bash", "makeRecentLanguagesScript.txt")
	err := cmd.Run()
	utils.HandleIfError(err)
	log.Println("Successfully created myRecentLanguages.svg")
}
