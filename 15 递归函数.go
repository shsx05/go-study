package main

// import "fmt"

// // 递归函数实现阶乘
// func factorial(n int) (result int) {
// 	if n > 0 {
// 		result = n * factorial(n-1)
// 		fmt.Println("当前进行的数字为", n)
// 		fmt.Println("当前结果为", result)
// 		return
// 	} else {
// 		return 1
// 	}
// }

// // 递归函数实现斐波那契数列
// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	// i := 5
// 	// res := factorial(i)
// 	// fmt.Printf("%d的阶乘为:%d", i, res)
// 	// 调用feibonacci函数
// 	c := make(chan int, 10)
// 	fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }
