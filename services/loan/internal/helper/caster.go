package helper

import "time"

func CastSQLTimeToTime(sqlTime string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", sqlTime)

	return t
}
