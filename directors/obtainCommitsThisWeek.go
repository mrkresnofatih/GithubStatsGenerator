package directors

import (
	"log"
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
)

func ObtainCommitsThisWeek() []int {
	// Get Top 30 Recent Repositories
	repos := ObtainRecentRepositories(25)

	// Instantiate Day-CommitCount Map
	weekCommits := map[string]int{}
	dates := utils.GetThisWeekDates()
	for _, value := range dates {
		weekCommits[value] = 0
	}
	result := []int{}

	// Top Day
	commitMax := 0

	for _, value := range repos {
		commits, err := apis.GetCommits(value.Name)
		utils.HandleIfError(err)

		for _, val := range *commits {
			date := utils.GetDateStringFromIso(val.Commit.Author.Date)
			log.Println("Getting Repo [" + value.Name + "] SHA [" + val.Sha + "] at [" + date + "]!")
			_, ok := weekCommits[date]
			if ok {
				weekCommits[date]++
				if weekCommits[date] > commitMax {
					commitMax = weekCommits[date]
				}
			}
		}
	}

	// Calculate x/100 PX
	for _, v := range dates {
		pixels := 25 + weekCommits[v]*75/commitMax
		result = append(result, pixels)
	}

	return result
}
