package directors

import (
	log "log"
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"sort"
	"time"
)

func ObtainTopLanguages(allRepos []apis.GetRepositoriesResponse) []LanguageUsage {
	sizes := map[string]int64{}
	total_size := int64(0)
	result := []LanguageUsage{}

	for _, value := range allRepos {
		if value.Language == "" {
			continue
		}

		_, ok := sizes[value.Language]
		actualSize, err := apis.GetLanguages(value.Name, value.Language)
		utils.HandleIfError(err)
		if ok {
			sizes[value.Language] += *actualSize
		} else {
			sizes[value.Language] = *actualSize
		}
		total_size += *actualSize

		log.Println("Obtained Actual Size of [" + value.Language + "] from Repo [" + value.Name + "]")
		time.Sleep(1 * time.Second)
	}

	for key, value := range sizes {
		val := float32(value*10000/total_size) / 100
		result = append(result, LanguageUsage{key, val})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Usage > result[j].Usage
	})

	if len(result) > 8 {
		result = result[:8]
	}

	return result
}
