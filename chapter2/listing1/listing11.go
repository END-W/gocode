package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// admin 需要一个 user 类型作为管理者，并附加权限
type admin struct {
	person user
	level  string
}

func main() {
	// 声明 user 类型的变量
	//var bill user

	// 结构字面量来完成初始化
	// 第一种方式
	// 声明 user 类型的变量，并初始化所有字段
	lisa := user{
		name:       "lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	// 第二种方式
	//lisa := user{"Lisa", "lisa@email.com", 123, true}

	// 声明 admin 类型的变量
	fred := admin{
		person: user{
			name:       "lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}

	fmt.Println("user: ", lisa)
	fmt.Println("admin: ", fred)
}
