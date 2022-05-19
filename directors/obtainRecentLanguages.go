package directors

import (
	"log"
)

func ObtainRecentLanguages() []string {
	repos := ObtainRecentRepositories(30)
	langs := []string{}
	langTypes := map[string]bool{}

	for _, value := range repos {
		_, ok := langTypes[value.Language]
		if !ok {
			langs = append(langs, value.Language)
			langTypes[value.Language] = true
			log.Println("Repo [" + value.Name + "] w/ Lang: " + value.Language)
		}
	}

	if len(langs) > 4 {
		langs = langs[:4]
	}

	return langs
}
