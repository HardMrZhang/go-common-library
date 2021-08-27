package strconv

import (
	"fmt"
	"strconv"
	"testing"
)

/**
strconv包实现了基本数据类型与其字符串表示的转换，
主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。
https://pkg.go.dev/strconv
*/
func Test_Strconv_01(t *testing.T) {
	/**
	Atoi()：函数用户将字符串类型的整数转为int类型
	*/
	p1 := "200"
	r, _ := strconv.Atoi(p1)
	fmt.Printf("type：%T value：%v \n", r, r)

}

func Test_Strconv_02(t *testing.T) {

	/**
	Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。

	ParseBoo: 它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

	ParseInt：返回字符串表示的整数值，接受正负号。

	ParseFloat：解析一个表示浮点数的字符串并返回其值。

	ParseUint：ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	*/
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-2", 10, 64)
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i, u)
}

func Test_Strconv_03(t *testing.T) {

	/**
	Format系列函数实现了将给定类型数据格式化为string类型数据的功能。

	func FormatBool(b bool) string：根据b值返回”true”或”false”。

	func FormatInt(i int64, base int) string：base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。

	func FormatUint(i uint64, base int) string：是FormatInt的无符号整数版本。

	func FormatFloat(f float64, fmt byte, prec, bitSize int) string：函数将浮点数表示为字符串并返回。
	fmt表示格式：’f’（-ddd.dddd）、’b’（-ddddp±ddd，指数为二进制）、’e’（-d.dddde±dd，十进制指数）、’E’（-d.ddddE±dd，十进制指数）、’g’（指数很大时用’e’格式，否则’f’格式）、’G’（指数很大时用’E’格式，否则’f’格式）。
	prec控制精度（排除指数部分）：对’f’、’e’、’E’，它表示小数点后的数字个数；对’g’、’G’，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。
	*/
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Println(s1, s2, s3, s4)

}
