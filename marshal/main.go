package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

// Class 班级
type Class struct {
	Title    string     `json:"title"`
	Students []*Student `json:"students"`
}

func main() {
	c := Class{
		Title:    "101",
		Students: make([]*Student, 0, 100),
	}
	for i := 0; i < 5; i++ {
		stu := &Student{
			ID:     i,
			Gender: "Man",
			Name:   fmt.Sprintf("stu%d", i),
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(&c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"}]}`
	c1 := Class{}
	err = json.Unmarshal([]byte(str), &c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}
