package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hinhchunhat struct {
	chieudai  float32 `desc:"Chieu dai hinh chu nhat"`
	chieurong float32 `desc:"Chieu rong hinh chu nhat"`
}

// Chu vi hinh chu nhat
//   - Cong thuc: (chieudai + chieurong) * 2
//
// @return float32
func (hcn *Hinhchunhat) chuvi() float32 {
	return (hcn.chieudai + hcn.chieurong) * 2
}

// Dien tich hinh chu nhat
//   - Cong thuc: chieudai * chieurong
//
// @return float32
func (hcn *Hinhchunhat) dientich() float32 {
	return hcn.chieudai * hcn.chieurong
}

// Lay du lieu tu ban phim
// - Sư dung Scanf
func inputHinhChuNhat() Hinhchunhat {
	var hinhchunhat Hinhchunhat

	for {
		fmt.Print("Vui long nhap kich thuoc chieu rong hinh chu nhat: ")
		_, err := fmt.Scanf("%f", &hinhchunhat.chieurong)
		if err == nil && hinhchunhat.chieurong > 0 {
			break
		}

		fmt.Println("⛔️ Kich thuoc chieu rong phai lon hon 0")
	}

	for {
		fmt.Print("Vui long nhap kich thuoc chieu dai hinh chu nhat: ")
		_, err := fmt.Scanf("%f", &hinhchunhat.chieudai)
		if err == nil && hinhchunhat.chieudai > 0 {
			break
		}

		fmt.Println("⛔️ Kich thuoc chieu chieu phai lon hon 0")
	}

	return hinhchunhat
}

// Lay du lieu tu ban phim
// 	- Sư dung Bufio
func inputHinhChuNhatV2() Hinhchunhat {
	var cr, cd float32

	for {
		var err error
		cr, err = readFloat("Vui long nhap kich thuoc chieu rong hinh chu nhat: ")
		if err != nil || cr <= 0 {
			fmt.Println("⛔️ Kich thuoc chieu rong phai lon hon 0")
		} else {
			break
		}
	}
	
	for {
		var err error
		cd, err = readFloat("Vui long nhap kich thuoc chieu dai hinh chu nhat: ")
		if err != nil || cd <= 0 {
			fmt.Println("⛔️ Kich thuoc chieu dai phai lon hon 0")
		} else {
			break
		}
	}
	
	return Hinhchunhat{
		chieurong: cr,
		chieudai: cd,
	}
}

func readFloat(prompt string) (float32, error) {
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
	numb, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}

	return float32(numb), nil
}

func main() {
	// hinhchunhat := Hinhchunhat{
	// 	chieudai: 20,
	// 	chieurong: 5,
	// }

	// fmt.Printf("Kich thuoc chu vi hinh chu nhat %.2f \n", hinhchunhat.chuvi())
	// fmt.Printf("Kich thuoc dien tich hinh chu nhat %.2f \n", hinhchunhat.dientich())

	hcn := inputHinhChuNhatV2()

	fmt.Printf("Kich thuoc chu vi hinh chu nhat %.2f \n", hcn.chuvi())
	fmt.Printf("Kich thuoc dien tich hinh chu nhat %.2f \n", hcn.dientich())

}
