package generics

import "fmt"

func filterTestToturial[T comparable](data []T, key T, f func(T) bool) bool {
	result := false
	for _, v := range data {
		if v == key {
			result = true
		}
	}
	return result
}

func RunFilter() {
	data := []int{1, 2, 3, 4, 5}
	key := 4
	fmt.Println("Original array int data need to filter: ", data)
	fmt.Println("Key need to filter: ", key)
	result := filterTestToturial(data, key, func(v int) bool {
		return v == key
	})
	if result {
		println("Found key ", key)
	} else {
		println("Not found key ", key)
	}

	data2 := []string{"a", "b", "c", "d", "e"}
	key2 := "x"
	fmt.Println("Original array string data need to filter: ", data2)
	fmt.Println("Key need to filter: ", key2)
	result2 := filterTestToturial(data2, key2, func(v string) bool {
		return v == key2
	})
	if result2 {
		println("Found key ", key2)
	} else {
		println("Not found key ", key2)
	}
}
