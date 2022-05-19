package utils

import (
	"math/rand"
	"time"
)

func GetColorPallete() []string {
	rand.Seed(time.Now().Unix())

	colors := []string{"#333C83", "#F24A72", "#00C897", "#F76E11", "#30AADD", "#5902EC", "#008E89", "#FE7E6D"}
	rand.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})

	return colors
}
