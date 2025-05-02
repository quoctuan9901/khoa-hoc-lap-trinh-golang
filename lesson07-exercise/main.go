package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tong(n int) int {
	// Dieu kien dung
	if n == 1 {
		return 1
	}

	// Goi de quy
	return n + tong(n-1)
}

func fibo(n int) {
	term1, term2 := 0, 1

	if n == 1 {
		fmt.Printf("Day so fibonacci: %d  \n", term1)
		fmt.Printf("Tong day so fibonacci: %d \n", term1)
	} else if n == 2 {
		fmt.Printf("Day so fibonacci: %d %d \n", term1, term2)
		fmt.Printf("Tong day so fibonacci: %d \n", term1+term2)
	} else {
		total := term1 + term2

		fmt.Printf("Day so fibonacci: %d %d ", term1, term2)

		for i := 3; i <= n; i++ {
			term3 := term1 + term2

			fmt.Printf("%d ", term3)

			total += term3

			term1, term2 = term2, term3
		}

		fmt.Println()

		fmt.Printf("Tong day so fibonacci: %d \n", total)
	}
}

func readInt(prompt string) (int, error) {
	// Thong bao nguoi dung nhap lieu
	fmt.Print(prompt)

	// Doc du lieu tu ban phim
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("error reading input: %v", err)
	}

	// Xoa bo khoang trang truoc va sau chuoi
	input = strings.TrimSpace(input)

	// Chuyen tu chuoi sang so
	numb, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}

	return numb, nil
}

func main() {

	for {
		fmt.Println()
		fmt.Println("-=-=-=-=- Menu Chuc Nang -=-=-=-=-")
		fmt.Println("[1] Tinh tong day so N phan tu")
		fmt.Println("[2] Hien thi va tinh tong day so Fibonacci")
		fmt.Println("[0] Thoat chuong trinh")
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

		var choice int
		for {
			var err error
			choice, err = readInt("Vui long nhap so chuc nang muon thuc hien: ")

			if err != nil || choice < 0 {
				fmt.Println("⛔️ Vui long nhap mot so nguyen")
			} else {
				break
			}
		}
		
		switch choice {
		case 1:
			var numb int

			for {
				var err error
				numb, err = readInt("Vui long nhap 1 so nguyen duong N: ")

				if err != nil || numb <= 0 {
					fmt.Println("⛔️ Vui long nhap vao la 1 so nguyen duong N")
				} else {
					break
				}
			}
			
			result := tong(numb)

			fmt.Printf("Tong tu 1 den %d co ket qua la: %d \n", numb, result)
		case 2:
			var numb int

			for {
				var err error
				numb, err = readInt("Vui long nhap so luong day so fibonacci muon hien thi: ")

				if err != nil || numb <= 0 {
					fmt.Println("⛔️ So luong day so fibonacci phai lon hon 0")
				} else {
					break
				}
			}
			
			fibo(numb)
		case 0:
			fmt.Println("🟢 Cam on da su dung chuong trinh. Bye !")
			return
		default:
			fmt.Println("🔴 Vui long chon dung so chuc nang muon thuc hien")
		}
	}
}
