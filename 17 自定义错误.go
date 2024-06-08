// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func division(a, b float64) (res float64, err error) {
// 	if b == 0 {
// 		err = errors.New("分母不能为0")
// 		return 0, err
// 	} else {
// 		res = a / b
// 		return res, nil
// 	}
// }
// func main() {
// 	// 使用语法糖
// 	// 1. errors包
// 	var text string
// 	text = "这是一个自定义错误"
// 	err1 := errors.New(text)
// 	fmt.Println(err1.Error())
// 	// 2. 使用fmt来定义
// 	err2 := fmt.Errorf("这是用fmt来定义的错误:%s", text)
// 	fmt.Println(err2.Error())
// 	// 使用var来定义
// 	var err error
// 	err = errors.New("这是使用var来定义的错误")
// 	fmt.Println(err.Error())
// 	//
// 	var num1, num2 float64
// 	fmt.Println("请输入两个数")
// 	fmt.Scan(&num1, &num2)
// 	result, err3 := division(num1, num2)
// 	if err3 != nil {
// 		fmt.Println(err3.Error())
// 	} else {
// 		fmt.Println("结果为: ", result)
// 	}
// }
