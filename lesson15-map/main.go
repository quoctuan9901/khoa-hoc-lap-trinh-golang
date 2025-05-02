package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	Role string
}

func main() {
	// employees := map[string]Employee{
	// 	"e1" : {"Tuan", 31, "Developer"},
	// 	"e2" : {"Teo", 35, "Manager"},
	// }

	// fmt.Printf("%+v \n", employees)

	// for key, value := range employees {
	// 	fmt.Printf("Gia tri cua key %s \n", key)
	// 	fmt.Printf("Name: %s \n", value.Name)
	// 	fmt.Printf("Age: %d \n", value.Age)
	// 	fmt.Printf("Role: %s \n", value.Role)
	// 	fmt.Println()
	// }

	studentSubject := map[string][]string{
		"Tuan": {"Toan", "Khoa hoc", "Python", "Golang"},
		"Teo": {"CNTT", "Khoa hoc"},
	}

	// fmt.Printf("%+v \n", studentSubject)

	// fmt.Printf("Mon hoc cua Tuan la: %s \n", studentSubject["Tuan"][0])
	// fmt.Printf("Mon hoc cua Tuan la: %s \n", studentSubject["Tuan"][1])

	for key, value := range studentSubject {
		for _, subject := range value {
			fmt.Printf("Mon hoc cua %s la: %s \n", key, subject)
		}
	}



	// drink := map[string]int{
	// 	"coffee": 500,
	// 	"tea": 300,
	// }

	// fmt.Printf("%+v \n", drink)

	// student := map[int]string{
	// 	10: "Tuan",
	// 	15: "Teo",
	// 	18: "Ti",
	// 	20: "John",
	// }

	// // fmt.Printf("%+v \n", student)

	// for k, v := range student {
	// 	fmt.Printf("Student[%d] co gia tri la: %s \n", k, v)
	// }

	// m := make(map[string]int)
	// m["a"] = 1
	// m["b"] = 2
	// m["c"] = 3

	// fmt.Printf("%+v \n", m)

	// monan := make(map[string]int)
	// monan["chao"] = 500
	// monan["com"] = 300

	// fmt.Printf("%+v \n", monan)

	// value, exists := monan["chao"]
	// if exists {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("Khong ton tai trong map")
	// }

}
