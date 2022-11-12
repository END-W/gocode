package main

import (
	_ "code/chapter1/matchers"
	"code/chapter1/search"
	"log"
	"os"
)

// init 在 main 之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}
