package directors

import (
	"log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
	"strconv"
)

func MakeRecentContributionsScript(data []int) {
	// Initialize the script
	f, err := os.Create("makeRecentContributionsScript.txt")
	utils.HandleIfError(err)
	defer f.Close()
	log.Println("Initialized makeRecentContributionsScript.txt")

	// CMD: copy template
	_, err2 := f.WriteString("cp templates/recentContributionsTemplate.txt myRecentContributions.svg \n")
	utils.HandleIfError(err2)
	log.Println("Created Copy Template CMD")

	// Get Color Pallette
	colorPallette := utils.GetColorPallete()

	for index, value := range data {
		color := colorPallette[index]
		pxl := strconv.Itoa(value)
		strIndex := strconv.Itoa(index)

		cmd := "sed -i 's/@@@DAY" + strIndex + "/" + pxl + "/' myRecentContributions.svg \n"
		_, er1 := f.WriteString(cmd)
		utils.HandleIfError(er1)

		cmdc := "sed -i 's/@@@COLOR" + strIndex + "/" + color + "/' myRecentContributions.svg \n"
		_, er2 := f.WriteString(cmdc)
		utils.HandleIfError(er2)

		log.Println("Created Replace CMD for @@@DAY & @@@COLOR")
	}
}
