package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给 admin 用户发送一个通知
	// 接口的嵌入的内部类型实现并没有提升到
	// 外部类型
	sendNotification(&ad)

	// 我们可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法没有被提升
	ad.notify()
}
