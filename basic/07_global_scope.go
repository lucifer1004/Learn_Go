package main

var a1 = "G"

func main() {
	n1()
	m1()
	n1()
}

func n1() {
	print(a1)
}

func m1() {
	a1 = "O" // 如果用 a1 := "O" 就是局部变量
	print(a1)
}