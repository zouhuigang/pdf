/*
go run calibre.go /mnt/hgfs/workspacego/src/vendor/github.com/zouhuigang/converter/example/gogs_zh/config.json
*/
package main

import (
	"os"

	"fmt"

	"github.com/zouhuigang/converter/converter"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("错误：缺少指定的json配置文件")
	} else {
		if converter, err := converter.NewConverter(args[0]); err != nil {
			fmt.Println(err.Error())
		} else {
			if err = converter.Convert(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
