package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user   // 嵌入的类型未公开
	Rights int
}
