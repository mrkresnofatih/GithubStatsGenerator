package directors

import (
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
	"strconv"
)

func MakeRecentLanguagesScript(data []string) {
	// Initialize the script
	f, err := os.Create("makeRecentLanguagesScript.txt")
	utils.HandleIfError(err)
	defer f.Close()
	log.Println("Initialized makeRecentLanguagesScript.txt")

	// CMD: copy template
	_, err2 := f.WriteString("cp templates/recentLanguagesTemplate.txt myRecentLanguages.svg \n")
	utils.HandleIfError(err2)
	log.Println("Created copy template CMD")

	for index, value := range data {
		cmdSed := "sed -i 's/@@@AVAILABLE" + strconv.Itoa(index) + "//' myRecentLanguages.svg \n"
		_, errr := f.WriteString(cmdSed)
		utils.HandleIfError(errr)

		log.Println("Created Replace CMD for @@@AVAILABLE" + strconv.Itoa(index))

		cmdSd := "sed -i 's/@@@LANG" + strconv.Itoa(index) + "/" + value + "/' myRecentLanguages.svg \n"
		_, errrr := f.WriteString(cmdSd)
		utils.HandleIfError(errrr)

		log.Println("Created Replace CMD for @@@LANG" + strconv.Itoa(index))
	}

	for i := 0; i < 3; i++ {
		cmd := "sed -i 's/@@@AVAILABLE" + strconv.Itoa(i) + "/opacity:0;/' myRecentRepositories.svg \n"
		_, er := f.WriteString(cmd)
		utils.HandleIfError(er)
		log.Println("Created Cleanup Template CMD for @@@AVAILABLE" + strconv.Itoa(i))
	}
}
