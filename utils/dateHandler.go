package utils

import (
	"strings"
	"time"
)

func GetDateStringFromIso(isoDate string) string {
	strs := strings.Split(isoDate, "T")
	return strs[0]
}

func GetThisWeekDates() []string {
	diffs := []int{1, 2, 3, 4, 5, 6, 7}
	dates := []string{}
	for _, value := range diffs {
		dt := time.Now().AddDate(0, 0, -1*value)
		strStamp := dt.Format(time.RFC3339)
		strDate := strings.Split(strStamp, "T")[0]
		dates = append(dates, strDate)
	}

	return dates
}
