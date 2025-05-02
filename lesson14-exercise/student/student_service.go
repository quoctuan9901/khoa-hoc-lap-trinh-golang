package student

import (
	"fmt"

	"quoctuan.com/school-management/utils"
)

var studentList []Student

func addStudent() {
	var scores []float64
	var id int
	fmt.Println("-=-=-=-=- Them Sinh Vien -=-=-=-=-")
	for {
		id = utils.GetPostiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, studentList) {
			break
		}

		fmt.Println("❌ Id da ton tai! Vui long nhap Id khac.")
	}

	name := utils.GetNonEmptyString("- Nhap ten: ")
	class := utils.GetNonEmptyString("- Nhap lop: ")
	totalPoint := utils.GetPostiveInt("- Nhap so luong diem: ")

	for i := 1; i <= totalPoint; i++ {
		score := utils.GetPostiveFloat(fmt.Sprintf("- Nhap diem %d: ", i))
		scores = append(scores, score)
	}

	student := Student{
		Id: id,
		Name: name,
		Class: class,
		Scores: scores,
	}

	studentList = append(studentList, student)

	fmt.Println("✅ Them sinh vien thanh cong!")
}

func deleteStudent() {
	fmt.Println("-=-=-=-=- Xoa Sinh Vien -=-=-=-=-")
	id := utils.GetPostiveInt("🗑️  Nhập ID sinh vien can xoa: ")
	for i, s := range studentList {
		if s.Id == id {
			studentList = append(studentList[:i], studentList[i+1:]...)
			fmt.Println("✅ Xoa sinh vien thanh cong!")
			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func updateStudent() {
	fmt.Println("-=-=-=-=- Sua Sinh Vien -=-=-=-=-")
	id := utils.GetPostiveInt("✏️  Nhập ID sinh vien can sua: ")

	for i, s := range studentList {
		if s.Id == id {
			fmt.Println("🔄 Nhap thong tin moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalString(fmt.Sprintf("- Nhap ten (%s): ", s.Name), s.Name)
			class := utils.GetOptionalString(fmt.Sprintf("- Nhap lop (%s): ", s.Class), s.Class)

			newScore := make([]float64, len(s.Scores))
			for idx, score := range s.Scores {
				temp := fmt.Sprintf("- Nhap diem %d (%.2f): ", idx + 1, score)
				newScore[idx] = utils.GetOptionalPostiveFloat(temp, score)
			}

			studentList[i] = Student{
				Id: id,
				Name: name,
				Class: class,
				Scores: newScore,
			}

			fmt.Println("✅ Cap nhat sinh vien thanh cong!")
			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func listStudent() {
	fmt.Println("-=-=-=-=- Danh Sach Sinh Vien -=-=-=-=-")
	if len(studentList) == 0 {
		fmt.Println("=> Khong co sinh vien nao trong danh sach.")
		return
	}

	for _, s := range studentList {
		fmt.Println(s.GetInfo())
	}
}

func searchStudent() {
	fmt.Println("-=-=-=-=- Tim Kiem Sinh Vien -=-=-=-=-")
	id := utils.GetPostiveInt("🔍  Nhap ID sinh vien can tim kiem:")
	for _, s := range studentList {
		if s.Id == id {
			fmt.Println("Tim thay sinh vien: ", s.GetInfo())
			return
		}
	}

	fmt.Println("❌ Khong tim thay sinh vien")
}

func StudentMenu() {
	for {
		utils.ClearScreen()

		fmt.Println("==== QUAN LY SINH VIEN ====")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPostiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			listStudent()
		case 5:
			searchStudent()
		case 6:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le!")
		}
	
		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
