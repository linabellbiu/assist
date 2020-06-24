package assist

import (
	"time"
)

func TimeFormatToUnix(t string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()
}

func UnixToTimeFormat(t int64) string {
	//获取本地location
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//时间戳转日期
	return time.Unix(t, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
}
