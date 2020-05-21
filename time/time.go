package time

import "time"

func FormatToUnix(t string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()
}
