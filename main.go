package main

import (
	executors "mrkresnofatih/githubstatsgenerator/executors"
)

func main() {
	executors.Authenticate()
	executors.CreateMyTopLanguages()
	executors.CreateMyRecentRepositories()
	executors.CreateMyRecentLanguages()
	executors.CreateMyRecentContributions()
	executors.UploadFilesToGithub()
	executors.CleanUp()
}
