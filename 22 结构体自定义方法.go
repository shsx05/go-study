package main

import "fmt"

// 类似于People类
type People struct {
	// People类的属性
	Name    string
	Age     int
	Address string
	Sex     string
}

// 类的方法
func (p *People) getInfo() string {
	return fmt.Sprintf("姓名: %s, 年龄: %d, 地址: %s, 性别: %s", p.Name, p.Age, p.Address, p.Sex)
}
func (p *People) Eat(food string) {
	fmt.Printf("%s吃了%s", p.Name, food)
}

func main() {
	var p1 People = People{
		Name:    "小红",
		Age:     12,
		Address: "北京",
		Sex:     "女",
	}
	info := p1.getInfo()
	fmt.Println(info)
	p1.Eat("宫保鸡丁")
}
