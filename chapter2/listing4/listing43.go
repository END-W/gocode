package main

import "fmt"

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 定义了程序里的管理员
type admin struct {
	name  string
	email string
}

// notify 使用指针接收者实现了 notifier 接口
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}

/*
方法集定义

Values                  Methods Receivers
-----------------------------------------------

	T                    (t T)
	*T                   (t T) and (t *T)

Methods 				Receivers Values
-----------------------------------------------

	(t T) 				  T and *T
	(t *T) 				  *T
*/
func main() {
	// 创建一个 user 类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}

	// 报错原因：编译器并不是总能自动获得一个值的地址
	//sendNotification(u)
	sendNotification(&u)
	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)

	// 展示接口多态性
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create an admin value and pass it to sendNotification.
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}
