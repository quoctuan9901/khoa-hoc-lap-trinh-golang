package array

import "fmt"

type Nhanvien struct {
	Id   int
	Name string
	Age  int
}

func Array() {
	employees := [...]Nhanvien{
		{Id: 1, Name: "Tuan", Age: 10},
		{Id: 2, Name: "An", Age: 15},
		{Id: 3, Name: "Binh", Age: 18},
		{Id: 4, Name: "Hoa", Age: 12},
	}

	// fmt.Printf("Id: %d \n", employees[1].Id)
	// fmt.Printf("Name: %s \n", employees[1].Name)
	// fmt.Printf("Age: %d \n", employees[1].Age)

	for _, value := range employees {
		fmt.Printf("Name: %s and Age: %d \n", value.Name, value.Age)
	}

	var number int
	fmt.Println(number)

	var numbers [5]int
	fmt.Println(numbers)

	// var character string
	// fmt.Println(character)

	// var characters [3]string
	// fmt.Println(characters)

	// var numbers [5]int
	// numbers[0] = 9  // 9 0 0  0 0
	// numbers[2] = 10 // 9 0 10 0 0
	// numbers[4] = 5  // 9 0 10 0 5
	// fmt.Println(numbers)

	// var numbers02 = [5]int{9, 10, 8}
	// numbers02[3] = 7
	// numbers02[0] = 1
	// fmt.Println(numbers02)

	// var numbers03 = [...]int{5, 8, 9, 10, 11}
	// fmt.Printf("Total array: %T \n", numbers03)
	// numbers03[4] = 9
	// fmt.Println(numbers03)

	// var numbers = [5]int{1,2,3,4,5}
	// fmt.Println(numbers[2])

	// var matrix = [2][3]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// }

	// fmt.Println(matrix)
	// fmt.Println(matrix[1][2])
	// matrix[1][1] = 9
	// matrix[0][2] = 10
	// fmt.Println(matrix)

	// numbers := [5]int{6,7,8,9,10}
	// fmt.Println(numbers[0])
	// fmt.Println(numbers[1])
	// fmt.Println(numbers[2])
	// fmt.Println(numbers[3])
	// fmt.Println(numbers[4])

	// fmt.Println(len(numbers))
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// numbers := [3][4]int{
	// 	{1,2,3,4},
	// 	{5,6,7,8},
	// 	{9,10,11,12},
	// }

	// for k := 0; k < len(numbers); k++ {
	// 	for j := 0; j < len(numbers[k]); j++ {
	// 		fmt.Println(numbers[k][j])
	// 	}
	// }

	// numbers := [5]int{6,7,8,9,10}
	// for key, val := range numbers {
	// 	fmt.Printf("Array numbers[%d] = %d \n", key, val)
	// }

	// numbers := [3][4]int{
	// 	{1,2,3,4},
	// 	{5,6,7,8},
	// 	{9,10,11,12},
	// }

	// for key1, val1 := range numbers {
	// 	for key2, val2 := range val1 {
	// 		fmt.Printf("Array numbers[%d][%d] = %d \n", key1, key2, val2)
	// 	}
	// }
}
