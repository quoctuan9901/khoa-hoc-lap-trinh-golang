package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Giangvien struct {
	Name   string `json:"hoten"`
	Gender int    `json:"gioitinh"`
	Phone  string `json:"dienthoai"`
}

func (gv *Giangvien) hienThiThongTin() {
	fmt.Printf("Ho ten: %s \n", gv.Name)
	fmt.Printf("Gioi tinh: %d \n", gv.Gender)
	fmt.Printf("Dien thoai: %s \n", gv.Phone)
}

func (gv *Giangvien) clear() {
	gv.Name = ""
	gv.Gender = 0
	gv.Phone = ""
}

func main() {

	quoctuan := Giangvien{
		Name:   "Vu Quoc Tuan",
		Gender: 1,
		Phone:  "0901234567",
	}
	// quoctuan.hienThiThongTin()

	// fmt.Println()

	// quoctuan.clear()

	// fmt.Println()

	// quoctuan.hienThiThongTin()

	output, err := json.Marshal(quoctuan)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(output))

}
