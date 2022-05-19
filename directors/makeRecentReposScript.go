package directors

import (
	"log"
	apis "mrkresnofatih/githubstatsgenerator/apis"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
	"strconv"
)

func MakeRecentReposScript(data []apis.GetRepositoriesResponse) {
	// Initialize the script
	f, err := os.Create("makeRecentRepositoriesScript.txt")
	utils.HandleIfError(err)
	defer f.Close()
	log.Println("Initialized makeRecentRepositoriesScript.txt")

	// CMD: copy template
	_, err2 := f.WriteString("cp templates/recentRepositoriesTemplate.txt myRecentRepositories.svg \n")
	utils.HandleIfError(err2)
	log.Println("Created Copy Template CMD")

	// Get Color Palette
	colorPallette := utils.GetColorPallete()

	for index, value := range data {
		color := colorPallette[index]

		temp := "<p class=\"pack\"><span class=\"repotag\" style=\"background:" + color + ";color:white;\"><code>" + value.Name + "<\\/code><\\/span><span class=\"rank\" style=\"background:whitesmoke;\">#" + strconv.Itoa(index) + "<\\/span><\\/p>"
		cmd := "sed -i 's/@@@REPO" + strconv.Itoa(index) + "/" + temp + "/' myRecentRepositories.svg \n"
		_, errr := f.WriteString(cmd)
		utils.HandleIfError(errr)

		log.Println("Created Replace CMD for @@@REPO" + strconv.Itoa(index))
	}

	for i := 0; i < 3; i++ {
		cmd := "sed -i 's/@@@REPO" + strconv.Itoa(i) + "//' myRecentRepositories.svg \n"
		_, er := f.WriteString(cmd)
		utils.HandleIfError(er)
		log.Println("Created Cleanup Template CMD for @@@REPO" + strconv.Itoa(i))
	}
}
