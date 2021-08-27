package pinyin

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"testing"
)

/**
汉语拼音转换工具。
*/
func Test_Pinyin(t *testing.T) {
	str := "中国人"
	//默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(str, a))
}
