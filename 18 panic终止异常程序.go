// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// func connectDatabase(address string, post int) (string, error) {
// 	if address == "" || post == 0 {
// 		return "", errors.New("数据库连接失败")
// 	} else {
// 		return "数据库连接成功", nil
// 	}
// }
// func main() {
// 	// 模拟连接数据库程序
// 	s, err := connectDatabase("", 0)
// 	for {
// 		time.Sleep(5 * time.Second)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			panic(err)
// 		} else {
// 			fmt.Println(s)
// 		}
// 	}
// }
