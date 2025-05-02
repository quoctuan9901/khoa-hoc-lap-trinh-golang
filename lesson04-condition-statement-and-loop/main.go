package main

import (
	"fmt"
)

func main() {

	// diem := 7

	// if diem >= 8 {
	// 	fmt.Println("Diem lon hon 8")
	// } else {
	// 	fmt.Println("Diem khong lon hon 8")
	// }

	// fmt.Println("Out !")

	// diem := -3

	// if diem > 8 { 							// 9, 10, ...
	// 	fmt.Println("Hoc sinh gioi")
	// } else if diem <= 8 && diem >= 6 {		// 6, 7, 8
	// 	fmt.Println("Hoc sinh kha")
	// } else if diem <= 5 && diem >= 3 {		// 3, 4, 5
	// 	fmt.Println("Hoc sinh trung binh")
	// } else {								// 0, 1, 2
	// 	fmt.Println("Hoc sinh yeu")
	// }

	// monan := "Bun"

	// switch monan {
	// case "Com", "Bun":
	// 	fmt.Println("Hom nay an com hoac an bun")
	// case "Chao", "Mi":
	// 	fmt.Println("Hom nay an chao hoac an mi")
	// case "Pho":
	// 	fmt.Println("Hom nay an pho")
	// default:
	// 	fmt.Println("Khong an gi ca")
	// }

	// Sô 1 đến 10: 1 2 3 4 5 6 7 8 9 10
	// i := 1
	// for ;i <= 10;  {
	// 	fmt.Printf("Number %d \n", i)
	// 	i++
	// }

	// Số 1 dến 100
	// 9 % 9 = 0
	// for i := 1; i <= 100; i += 2 {
	// 	if i % 9 == 0 {
	// 		fmt.Printf("Number %d \n", i)
	// 	}	
	// }

	// break
	for i := 1; i <= 10; i++ {
		if (i % 2 == 0) {
			break;
		}

		fmt.Printf("Number %d \n", i)
	}

	fmt.Println("End code")

	// continue
	for i := 1; i <= 10; i++ {
		if (i % 2 == 0) {
			continue;
		}

		fmt.Printf("Number %d \n", i)
	}

	fmt.Println("End code")
}
