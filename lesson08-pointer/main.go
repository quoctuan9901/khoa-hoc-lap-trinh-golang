package main

import "fmt"

func lythuyetPointer() {
	name := "Vu Quoc Tuan"
	fmt.Println("-=-=-=-=- Information name variable -=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()

	// Tạo con tro (pointer)
	ptrName := &name
	fmt.Println("-=-=-=-=- Information ptrName variable -=-=-=-=-")
	fmt.Printf("Data type: %T \n", ptrName)
	fmt.Printf("Value: %v \n", ptrName)
	fmt.Printf("Address: %v \n", &ptrName)

	fmt.Printf("Find value name from ptrName: %v \n", *ptrName)

	fmt.Println()
	// Tao con tro
	ptrName2 := &ptrName
	fmt.Println("-=-=-=-=- Information ptrName2 variable -=-=-=-=-")
	fmt.Printf("Data type: %T \n", ptrName2)
	fmt.Printf("Value: %v \n", ptrName2)
	fmt.Printf("Address: %v \n", &ptrName2)

	fmt.Printf("Find value ptrName from ptrName2: %v \n", *ptrName2)
	fmt.Printf("Find value name from ptrName2: %v \n", **ptrName2)
}

// &name
func updateName(name *string) {
	*name = "Nguyen Van A"

	fmt.Println("-=-=-=-=- Information name variable inside updateName -=-=-=-=-")
	fmt.Printf("Value: %s \n", *name)
}

func main() {
	name := "Vu Quoc Tuan"
	fmt.Println("-=-=-=-=- Information name variable before run updateName -=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

	fmt.Println()
	updateName(&name)
	fmt.Println()

	fmt.Println("-=-=-=-=- Information name variable after run updateName -=-=-=-=-")
	fmt.Printf("Data type: %T \n", name)
	fmt.Printf("Value: %v \n", name)
	fmt.Printf("Address: %v \n", &name)

}
