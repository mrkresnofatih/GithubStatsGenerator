package apis

import (
	utils "mrkresnofatih/githubstatsgenerator/utils"
)

func GetNewestCommit(repo string) (*GetCommitResponse, error) {
	commits := []GetCommitResponse{}
	resp, err := GetCommits(repo)
	utils.HandleIfError(err)

	if len(*resp) == 0 {
		e := utils.BuildError("Empty Repository")
		return nil, e
	}
	commits = append(commits, *resp...)
	return &commits[0], nil
}
