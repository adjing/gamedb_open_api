/*
> go mod init syscom
*/
package sys

import (
	"fmt"
	"time"
)

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func GetBeiJingTime() string {
	t, _ := TimeIn(time.Now(), "Asia/Shanghai")
	return fmt.Sprintf("%s", t.Format("2006-01-02 15:04:05"))
}

func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

func GetTime_Now_Year() int {
	return time.Now().Year()
}

func GetTime_Now_Month() int {
	return int(time.Now().Month())
}

func GetTime_Now_Day() int {
	return time.Now().Day()
}

func GetTime_Now_Hour() int {
	return time.Now().Hour()
}

func GetTime_Now_Minute() int {
	return time.Now().Minute()
}

//
func GetTime_Yesterday_Year() int {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Year()
}

func GetTime_Yesterday_Month() int {
	yesterday := time.Now().AddDate(0, 0, -1)
	return int(yesterday.Month())
}

func GetTime_Yesterday_Day() int {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Day()
}

//tomorrow

//中间下划线间隔
func GetTime_Yesterday_YMD() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	//
	var txt = fmt.Sprintf("%d_%d_%d", yesterday.Year(), int(yesterday.Month()), yesterday.Day())
	return txt
}

// 例 GetTimestamp("2018-06-07 12:00:00")
func GetTimestamp(t string) (err error, timestamp int64) {
	var formatTime time.Time
	if formatTime, err = time.Parse(t, "2018-06-07 12:00:00"); err != nil {
		timestamp = 0
		return
	}

	timestamp = formatTime.Unix()
	return
}

// 例 GetAfterNowTimestamp(0,)
func GetAfterNowTimestamp(hour, minute, second time.Duration) int64 {
	return time.Now().Add(hour*time.Hour + minute*time.Minute + second*time.Second).Unix()
}
