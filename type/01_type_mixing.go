package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 15
	//b = a + a // 编译错误
	b = int32(a) + int32(a)
	b = b + 5 // 因为 5 是常量，所以可以通过编译
	fmt.Println(b)
}
