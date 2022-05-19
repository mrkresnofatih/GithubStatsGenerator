package directors

import (
	log "log"
	utils "mrkresnofatih/githubstatsgenerator/utils"
	"os"
	"strconv"
)

func MakeTopLanguagesScript(data []LanguageUsage) {
	// Initialize the script
	f, err := os.Create("makeTopLanguagesScript.txt")
	utils.HandleIfError(err)
	defer f.Close()
	log.Println("Initialized makeTopLanguagesScript.txt File!")

	// CMD: copy template
	_, err2 := f.WriteString("cp templates/topLanguagesTemplate.txt myTopLanguages.svg \n")
	utils.HandleIfError(err2)
	log.Println("Created Copy Template CMD")

	// Get Color Pallette
	colorPallette := utils.GetColorPallete()

	for index, value := range data {
		color := colorPallette[index]
		usg := strconv.FormatFloat(float64(value.Usage), 'f', -1, 32)

		temp := "<p class=\"language\"><span class=\"lang\" style=\"background:" + color + ";color:white;\">" + value.Language + "<\\/span> <span class=\"percent\" style=\"background:whitesmoke;\">" + usg + "%<\\/span><\\/p>"

		cmd := "sed -i 's/@@@LANG" + strconv.Itoa(index) + "/" + temp + "/' myTopLanguages.svg \n"
		_, errr := f.WriteString(cmd)
		utils.HandleIfError(errr)

		log.Println("Created Replace CMD for @@@LANG" + strconv.Itoa(index))
	}

	for i := 0; i < 8; i++ {
		cmd := "sed -i 's/@@@LANG" + strconv.Itoa(i) + "//' myTopLanguages.svg \n"
		_, er := f.WriteString(cmd)
		utils.HandleIfError(er)
		log.Println("Created Cleanup Template CMD for @@@LANG" + strconv.Itoa(i))
	}
}
