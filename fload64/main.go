package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := 2.2232149942234
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	fmt.Println(value)
}
