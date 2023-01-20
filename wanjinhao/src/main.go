package main

import "fmt"

func main() {
	slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 1, 2) // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(len(slice2), cap(slice2))

}
