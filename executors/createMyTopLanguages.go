package executors

import (
	log "log"
	directors "mrkresnofatih/githubstatsgenerator/directors"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	exec "os/exec"
)

func CreateMyTopLanguages() {
	log.Println("Starting Process: Create myTopLanguages.svg")
	allRepos := directors.ObtainAllRepositories()
	repos := directors.ObtainTopLanguages(allRepos)
	directors.MakeTopLanguagesScript(repos)

	cmd := exec.Command("bash", "makeTopLanguagesScript.txt")
	err := cmd.Run()
	utils.HandleIfError(err)
	log.Println("Successfully created myTopLanguages.svg")
}
