package file

import (
	"fmt"
	"github.com/hpcloud/tail"
	"testing"
	"time"
)

/**
实时读取文件内容
*/

func Test_File(t *testing.T) {
	fileName := "./test.txt"
	config := tail.Config{
		//从文件哪里读
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,  //重新打开
		MustExist: false, //文件如果不存在就报错
		Poll:      false, //轮询文件更改，而不是使用inotify
		Follow:    true,  //是否跟随
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("跟踪文件失败", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail 文件关闭打开,filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}

}
