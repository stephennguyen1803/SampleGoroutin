package generics

import "fmt"

func mapTestToturial[T, K comparable](data []T, f func(T) K) []K {
	var result []K
	for _, v := range data {
		result = append(result, f(v))
	}
	return result
}

func RunMap() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array int data: ", data)
	result := mapTestToturial(data, func(v int) int {
		return v * 2
	})
	fmt.Println(result)

	data2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Original array string data: ", data2)
	resutl2 := mapTestToturial(data2, func(v string) string {
		return "hello " + v
	})

	fmt.Println(resutl2)
}
