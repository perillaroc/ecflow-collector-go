package node_checker

import (
	"fmt"
	"time"
)

func ParseClockUTC(s string) (time.Time, error) {
	currentTime := time.Now().UTC()
	utcDate := currentTime.Format("2006-01-02")
	dateString := fmt.Sprintf("%s %s UTC", utcDate, s)

	result, err := time.Parse("2006-01-02 15:04 MST", dateString)
	return result, err
}
