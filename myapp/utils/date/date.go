package date

import (
	"time"
)

func GetDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}
