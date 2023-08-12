package MyMap

import "fmt"

func Demo() {
	var m1 map[string]int
	m1 = make(map[string]int)
	fmt.Println("m1=", m1, " len=", len(m1))
	m1["a"] = 1
	fmt.Println("m1=", m1, " len=", len(m1))

	m1["b"] = 2
	fmt.Println("m1=", m1, " len=", len(m1))

	m1["a"] = 3
	fmt.Println("m1=", m1, " len=", len(m1))

	delete(m1, "b")
	fmt.Println("m1=", m1, " len=", len(m1))

	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("m2=", m2, " len=", len(m2))

	var m3 = map[string]int{"a": 1, "b": 2}
	fmt.Println("m3=", m3, " len=", len(m3))

	value, ok := m3["test"]
	if ok {
		fmt.Println("value=", value)
	} else {
		fmt.Println("not found")
	}

	for key, value := range m3 {
		fmt.Println("key=", key, " value=", value)
	}
}
