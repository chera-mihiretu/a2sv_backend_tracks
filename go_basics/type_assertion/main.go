package main

import "fmt"

func main() {

	print("chera")

	m_map := make(map[string]string)

	m_map["Chera"] = "Is Cool"

	m_map["Carra"] = "Is not that much cool"

	print(m_map)

	print(20.34)

}

func print(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer type")
	case map[string]string:
		fmt.Println("This is json")
	default:
		fmt.Println("WTF is this")
	}
	fmt.Println(i)
}
