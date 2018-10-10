package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5, 6}
	sum := 0

	for _, v := range nums {
		sum += v
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {

		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
