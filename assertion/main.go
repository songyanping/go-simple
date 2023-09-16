package main

import "fmt"

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	case Pasen:
		fmt.Println(x, "is Pasen")
	default:
		fmt.Println(x, "not type matched")
	}
}

// 断言示例
func main() {
	findType(10)      // int
	findType("hello") // string

	var k interface{} // nil
	findType(k)

	findType(10.23) //float64

	p := Pasen{
		Name: "Jack",
		Age:  30,
	}
	findType(p)
}

type Pasen struct {
	Name string
	Age  int
}
