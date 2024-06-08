package main

// import "fmt"

// func main() {
// 	// 创建对象的第一种方法
// 	// 利用语法糖
// 	// teacherAge := make(map[string]int)
// 	// // 赋值
// 	// teacherAge["张三"] = 18
// 	// teacherAge["李四"] = 19
// 	// teacherAge["王五"] = 20
// 	// fmt.Println("teacherAge:", teacherAge)
// 	// // 也可以直接赋值
// 	// teacherAge2 := map[string]int{
// 	// 	"张三": 18,
// 	// 	"李四": 19,
// 	// 	"王五": 20,
// 	// }
// 	// fmt.Println("teacherAge2:", teacherAge2)
// 	// // 也可以使用var声明
// 	// var teacheraddress map[string]string
// 	// // 当你用var来声明对象时，必须初始化
// 	// teacheraddress = make(map[string]string)
// 	// teacheraddress["张三"] = "北京"
// 	// teacheraddress["李四"] = "上海"
// 	// teacheraddress["王五"] = "广州"
// 	// fmt.Println("teacheraddress:", teacheraddress)

// 	// // 获取对象的值
// 	// fmt.Println("张三的地址:", teacheraddress["张三"])
// 	// // 通过变量来取值
// 	// teachename := "李四"
// 	// fmt.Printf("查询%s的地址为%s\n", teachename, teacheraddress[teachename])
// 	// // 通过循环来取值
// 	// for k, v := range teacheraddress {
// 	// 	fmt.Printf("老师：%s, 地址：%s\n", k, v)
// 	// }
// 	// // 删除值 -- delete
// 	// delete(teacheraddress, "张三")
// 	// fmt.Println("修改后的值", teacheraddress)

// 	// var studentName map[int]string
// 	// studentName = make(map[int]string)
// 	// studentName[1] = "小明"
// 	// studentName[2] = "小红"
// 	// studentName[3] = "小雷"
// 	// fmt.Println("初始值: ", studentName)
// 	// studentDeleteId := 1
// 	// delete(studentName, studentDeleteId)
// 	// fmt.Println("修改后的数值", studentName)
// }
