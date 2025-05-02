package main

import "fmt"

func main() {
	// fmt.Print("Hello Tuan.")
	// fmt.Print("Goodbye Tuan.")

	// fmt.Println("Hello Tuan.")
	// fmt.Println("Goodbye Tuan.")

	// var fullName = "Vu Quoc Tuan"
	// var age = 35
	// fmt.Printf("Hello %s and your age is %d. \n", fullName, age)

	// var firstName, lastName string
	// fmt.Print("Please enter your name: ")
	// fmt.Scan(&firstName, &lastName)
	// fmt.Println("My name is ", firstName, lastName)

	// var firstName, lastName string
	// fmt.Print("Please enter your firstname: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Please enter your lastname: ")
	// fmt.Scanln(&lastName)
	// fmt.Println("My name is ", firstName, lastName)

	// var name string
	// var age int
	// fmt.Println("Please enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Please enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("My name is %s and age is %d \n", name , age)

	// message := fmt.Sprint("My name is ", "Tuan")
	// fmt.Println(message)

	// message := fmt.Sprintln("My name is ", "Tuan")
	// fmt.Println(message)

	// name := "Tuan"
	// age := 35
	// message := fmt.Sprintf("My name is %s and age is %d", name, age)
	// fmt.Println(message)

	ten := "Vu Quoc Tuan"
	tuoi := 35
	chieucao := 1.75
	daTotNghiep := true
	phanTram := 10

	fmt.Printf("Kieu du lieu cua bien ten: %T \n", ten)
	fmt.Printf("Kieu du lieu cua bien tuoi: %T \n", tuoi)
	fmt.Printf("Kieu du lieu cua bien chieucao: %T \n", chieucao)
	fmt.Printf("Kieu du lieu cua bien daTotNghiep: %T \n", daTotNghiep)
	fmt.Printf("Kieu du lieu cua bien phanTram: %T \n", phanTram)

	fmt.Printf("Toi ten la: %v \n", ten)

	fmt.Printf("Toi ten la %s va toi %d tuoi \n", ten, tuoi)

	fmt.Printf("Chieu cao cua toi la %.2fm \n", chieucao)
	fmt.Printf("Chieu cao cua toi la %.5fm \n", chieucao)
	fmt.Printf("Chieu cao cua toi la %.1fm \n", chieucao)

	fmt.Printf("Toi da tot nghiep: %t \n", daTotNghiep)

	fmt.Printf("Toi da hoc duoc %d%% \n", phanTram)
}