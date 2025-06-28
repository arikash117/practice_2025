package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var intVar int = 0
	var stringVar string = ""
	var boolean bool = true
	var f32 float32 = 0
	var f64 float64 = 0

	fmt.Println("Sizes of types in bytes:")
	fmt.Println("int:", unsafe.Sizeof(intVar))
	fmt.Println("string:", unsafe.Sizeof(stringVar))
	fmt.Println("boolean:", unsafe.Sizeof(boolean))
	fmt.Println("float32:", unsafe.Sizeof(f32))
	fmt.Println("float64:", unsafe.Sizeof(f64))
}
