package cmd

import (
	"strings"
	"time"
)

var (
	todays    = []string{"today", "今日", "now", "今"}
	tomorrows = []string{"tomorrow", "明日"}
)

func containsString(a []string, target string) bool {
	for _, r := range a {
		if target == r {
			return true
		}
	}
	return false
}

func parseDate(date string) (time.Time, error) {
	now := time.Now()
	date = strings.ReplaceAll(date, "/", "-")
	if containsString(todays, date) {
		return now, nil
	} else if containsString(tomorrows, date) {
		return now.AddDate(0, 0, 1), nil
	}
	d, err := time.ParseInLocation(dateFormat, date, now.Location())
	if err != nil {
		return now, err
	}

	return d, nil
}
