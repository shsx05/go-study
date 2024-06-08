// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func connectDatabase(address string, post int) (string, error) {
// 	if address == "" || post == 0 {
// 		return "", errors.New("数据库连接失败")
// 	} else {
// 		return "数据库连接成功", nil
// 	}
// }
// func returnDataToFrontend(msg string) {
// 	fmt.Println(msg)
// }

// func main() {
// 	msg := "返回前端的数据"
// 	defer returnDataToFrontend(msg)
// 	_, err := connectDatabase("", 0)
// 	if err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// }
