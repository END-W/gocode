package main

import (
	"code/chapter2/listing6/entities"
	"fmt"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	// 设置未公开的内部类型的
	// 公开字段的值
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
