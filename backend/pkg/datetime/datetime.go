package datetime

import "time"

var wib *time.Location

func init() {
	wib, _ = time.LoadLocation("Asia/Jakarta")
}

func ParseDateWIB(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", date, wib)
}
