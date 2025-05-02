package teacher

import (
	"fmt"

	"quoctuan.com/school-management/utils"
)

var teacherList []Teacher

func addTeacher() {
	var id int
	fmt.Println("-=-=-=-=- Them Giang Vien -=-=-=-=-")
	for {
		id = utils.GetPostiveInt("- Nhap id: ")
		if utils.IsIdUnique(id, teacherList) {
			break
		}

		fmt.Println("❌ Id da ton tai! Vui long nhap Id khac.")
	}

	name := utils.GetNonEmptyString("- Nhap ten: ")
	subject := utils.GetNonEmptyString("- Nhap mon giang day: ")
	baseSalary := utils.GetPostiveFloat("- Nhap luong co ban: ")
	bonus := utils.GetPostiveFloat("- Nhap thuong: ")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}

	teacherList = append(teacherList, teacher)

	fmt.Println("✅ Them giang vien thanh cong!")
}

func deleteTeacher() {
	fmt.Println("-=-=-=-=- Xoa Giang Vien -=-=-=-=-")
	id := utils.GetPostiveInt("🗑️  Nhập ID giang vien can xoa: ")
	for i, t := range teacherList {
		if t.Id == id {
			teacherList = append(teacherList[:i], teacherList[i+1:]...)
			fmt.Println("✅ Xoa giang vien thanh cong!")
			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func updateTeacher() {
	fmt.Println("-=-=-=-=- Sua Giang Vien -=-=-=-=-")
	id := utils.GetPostiveInt("✏️  Nhập ID giang vien can sua: ")

	for i, t := range teacherList {
		if t.Id == id {
			fmt.Println("🔄 Nhap thong tin moi (Nhan Enter de giu nguyen gia tri hien tai)")
			name := utils.GetOptionalString(fmt.Sprintf("- Nhap ten (%s): ", t.Name), t.Name)
			subject := utils.GetOptionalString(fmt.Sprintf("- Nhap mon giang day (%s): ", t.Subject), t.Subject)
			baseSalary := utils.GetOptionalPostiveFloat(fmt.Sprintf("- Nhap luong co bang (%.2f): ", t.BaseSalary), t.BaseSalary)
			bonus := utils.GetOptionalPostiveFloat(fmt.Sprintf("- Nhap thuong (%.2f): ", t.Bonus), t.Bonus)

			teacherList[i] = Teacher{
				Id: id,
				Name: name,
				Subject: subject,
				BaseSalary: baseSalary,
				Bonus: bonus,
			}

			fmt.Println("✅ Cap nhat giang vien thanh cong!")
			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func listTeacher() {
	fmt.Println("-=-=-=-=- Danh Sach Giang Vien -=-=-=-=-")
	if len(teacherList) == 0 {
		fmt.Println("=> Khong co giang vien nao trong danh sach.")
		return
	}

	for _, t := range teacherList {
		fmt.Println(t.GetInfo())
	}
}

func searchTeacher() {
	fmt.Println("-=-=-=-=- Tim Kiem Giang Vien -=-=-=-=-")
	id := utils.GetPostiveInt("🔍  Nhap ID giang vien can tim kiem:")
	for _, t := range teacherList {
		if t.Id == id {
			fmt.Println("Tim thay giang vien: ", t.GetInfo())
			return
		}
	}

	fmt.Println("❌ Khong tim thay giang vien")
}

func TeacherMenu() {
	for {
		utils.ClearScreen()

		fmt.Println("==== QUAN LY GIANG VIEN ====")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Quay lai")

		choice := utils.GetPostiveInt("👉 Chon chuc nang: ")

		switch choice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			updateTeacher()
		case 4:
			listTeacher()
		case 5:
			searchTeacher()
		case 6:
			return
		default:
			fmt.Println("❌ Lua chon khong hop le!")
		}

		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
