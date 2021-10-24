package main

import "fmt"

type person struct {
	name string
	age int32
}
type student struct {
	personBasicMSG person
	grade float64
}
type teacher struct {
	personBasicMAG person
	salary float64
}
func main() {
	stu:=student{
		personBasicMSG: person{
			age: 12,
			name: "刘海斌",
		},
		grade: 613.5,
	}
	tea:=teacher{personBasicMAG: person{
		age: 38,
		name: "lbz",
		},
		salary: 6610.5,
	}
	findType(stu)
	findType(tea)
	findType(int32(21))
}
func findType(v interface{}) {
	//使用接口断言实现函数使用使的得心应手（参数的传入没有类型的限制）：
	//接口断言常常配合空接口使用，这样任意类型都可以检测到
	switch v.(type) {
	case student:
		fmt.Println("student类型",v.(student))
	case teacher:
		fmt.Println("teacher类型",v.(teacher))
	case int32:
		fmt.Println("int32")
	default:
		fmt.Println("anther type")
	}
}
