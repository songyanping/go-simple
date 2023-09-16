package main

import "fmt"

func main() {

	//getAlertLevel("INFO")
	demo()

}

func getAlertLevel(level string) string {
	m := map[string]string{
		"INFO":     "P3",
		"WARN":     "P2",
		"CRITICAL": "P1",
	}
	value, ok := m[level]
	if ok {
		return value
	}
	return "default"
}

func demo() {

	var list []float64
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	var a float64
	for _, i := range list {
		a = a + i
	}
	fmt.Println(a)
}
