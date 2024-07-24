// 练习：斐波纳契闭包
// 让我们用函数做些好玩的。

// 实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 (0, 1, 1, 2, 3, 5, ...)。

package main

import "fmt"

// fibonacci 是返回一个「返回一个 int 的函数」的函数
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, a+b
		return c
	}
}

func main_fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
