package main

import (
	"fmt"
	"strconv"
)

func main() {
	/** Bai 01 **/
	// var xhtml = ""
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}

	// 	xhtml += strconv.Itoa(i)

	// 	if i != 100 {
	// 		xhtml += ","
	// 	} else {
	// 		xhtml += "\n"
	// 	}
	// }

	// fmt.Print(xhtml)

	/** Bai 02 **/
	// count := 0
	// xhtml := ""
	// for i := 1; i <= 100; i++ {
	// 	if i % 2 == 1 {
	// 		if count > 0 {
	// 			xhtml += ","
	// 		}

	// 		xhtml += strconv.Itoa(i)

	// 		count++

	// 		if count == 3 {
	// 			xhtml += "\n"
	// 			count = 0
	// 		}
	// 	}
	// }

	// if count > 0 {
	// 	xhtml += "\n"
	// }

	// fmt.Print(xhtml)

	/** Bai 03 **/
	var batdau, ketthuc int
	

	fmt.Print("Vui long nhap so bat dau: ")
	fmt.Scanf("%d", &batdau)

	fmt.Print("Vui long nhap so ket thuc: ")
	fmt.Scanf("%d", &ketthuc)

	if batdau == 0 || ketthuc == 0 {
		fmt.Println("⛔️ So bat dau va so ket thuc phai lon hon 0")
	} else if ketthuc < batdau {
		fmt.Println("⛔️ So ket thuc phai lon hon so bat dau")
	} else {
		var xhtml = ""

		for k := batdau; k <= ketthuc; k++ {
			xhtml += "Bang cuu chuong " + strconv.Itoa(k) + "\n"

			for l := 1; l <= 10; l++ {
				xhtml += strconv.Itoa(k) + " x " + strconv.Itoa(l) + " = " + strconv.Itoa(k * l) + "\n"
			}
		}

		fmt.Print(xhtml)
		
	}

}
