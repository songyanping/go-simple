package main

import (
	"fmt"
	"strconv"
	"time"
)

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
//var timeLayoutStr = 2006/01/02 03:04:05 //合法, 格式可以改变
//var timeLayoutStr = 2019/01/02 15:04:05 //不合法, 时间必须是2016年1月2号这个时间

func timeFormat() {
	t := time.Now() //当前时间
	fmt.Println("time.Now输出:", t)

	tStamp := t.Unix() //时间戳
	fmt.Println("t.Unit输出:", tStamp)

	ts := t.Format(timeLayoutStr) //time转string
	fmt.Printf("字符串类型的时间:%s\n", ts)

	st, _ := time.Parse(timeLayoutStr, ts) //string转time
	fmt.Printf("时间类型的时间:%s", st)

	//在go中, 可以格式化一个带前后缀的时间字符串
	prefixTStr := "PREFIX-- 2019-01-01 -TEST- 10:31:12 --SUFFIX"       //带前后缀的时间字符串
	preTimeLayoutStr := "PREFIX-- 2006-01-02 -TEST- 15:04:05 --SUFFIX" //需要转换的时间格式, 格式和前后缀需要一致, 这种写法的限制很大, 但一些特殊场景可以用到
	prefixTime, _ := time.Parse(preTimeLayoutStr, prefixTStr)
	fmt.Println(prefixTime)

	//时间加减 time.ParseDuration()
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	at, _ := time.ParseDuration("2h") //2个小时后的时间, 负数就是之前的时间
	fmt.Println((t.Add(at)).Format(timeLayoutStr))

	//两个时间差
	sub := t.Sub(prefixTime)
	fmt.Println(sub.Seconds()) //秒,  sub.Minutes()分钟,  sub.Hours()小时...

}

func UTC8() {
	// 获取当前时间
	currentTime := time.Now()
	// 添加8小时的持续时间
	utcTime := currentTime.Add(8 * time.Hour)
	// 将结果转换为 UTC 时间戳（秒数）
	utcTimestamp := utcTime.Unix()
	fmt.Println("UTC 时间戳（加8小时）:", utcTimestamp)
}

// 毫秒(13位)时间戳转成字符串格式化的时间
func timeStamptoStr() {

	//1.毫秒时间戳转成字符串格式化的时间
	s := "1685646149981"
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(i)
	t1 := time.UnixMilli(i)
	fmt.Println(t1)

	ts := t1.Format("2006-01-02 15:04:05") //time转string
	fmt.Println(ts)
}

// 字符串时间转成毫秒时间戳
func strToTimeStamp() {

	//1.字符串时间转成毫秒时间戳
	layout := "2006-01-02 15:04:05"
	str := "2023-08-09 13:16:00"

	// 使用 Parse 将时间字符串转换为 Time 类型
	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 Unix 函数将 Time 类型转换为毫秒时间戳
	//timestamp := t.Unix()
	timestamp := t.UnixMilli()
	fmt.Println(timestamp)
}

func main() {
	strToTimeStamp()

}
