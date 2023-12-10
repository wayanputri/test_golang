package helper

import (
	"fmt"
	"time"
)

func ParseTime(str string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}
	}
	return parsedTime
}
