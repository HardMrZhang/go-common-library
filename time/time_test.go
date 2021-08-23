package time

import (
	"fmt"
	"testing"
	"time"
)

/**
time包提供了时间的显示和测量用的函数。日历的计算采用的是公历。
*/
func Test_Time_01(t *testing.T) {
	/**
	time.Time类型表示时间,time.Now()函数获取当前的时间对象
	然后可以获取时间对象的年月日时分秒信息
	*/
	now := time.Now()
	fmt.Printf("当前时间为：%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func Test_Time_02(t *testing.T) {
	/**
	时间戳是至1970年1月1日至当前时间的总毫秒数
	*/
	now := time.Now()
	fmt.Printf("当前时间为：%v\n", now)
	timestamp1 := now.Unix()     //毫秒时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳

	fmt.Printf("当前时间戳：%v\n", timestamp1)
	fmt.Printf("当前时间戳：%v\n", timestamp2)

	//使用time.Unix()函数可用将时间戳转为时间格式
	timeObj := time.Unix(timestamp1, 0)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func Test_Time_03(t *testing.T) {
	/**
	time.Duration是time包定义的一个类型,他代表两个时间点之间经过的时间,以纳秒为单位
	time.Duration表示一段时间间隔,最长时间段大约290年
	*/
	/**
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	*/
}

func Test_Time_04(t *testing.T) {
	//当前时间+时间间隔,比如一个小时后的时间
	//add
	now := time.Now()
	later := now.Add(time.Hour)
	fmt.Println(later)

	//sub：求两个时间的差值,返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。
	//Equal:判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
	//Before：如果t代表的时间点在u之前，返回真；否则返回假。
	//After：如果t代表的时间点在u之后，返回真；否则返回假。
}

func Test_Time_05(t *testing.T) {
	/**
	go语言的格式化时间模板不是常见的y-m-d H:m:s,而是使用2006年1月2号15点04分
	*/
	now := time.Now()
	//24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//12小时制
	fmt.Println(now.Format("2006-01-02 3:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func Test_Time_06(t *testing.T) {
	/**
	解析字符串格式的时间
	*/
	now := time.Now()
	fmt.Println(now)

	//设置时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
