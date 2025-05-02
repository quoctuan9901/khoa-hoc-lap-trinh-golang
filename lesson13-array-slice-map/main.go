package main

func main() {

	// Khởi tạo slice
	// var numbers []int
	// fmt.Println(numbers)

	// [] => Slice rỗng

	// Slice => 0 có kích thước
	// slice := []int{1,2,3,4,5}
	// fmt.Println(slice)

	// Array => có kích thước
	// array := [5]int{1,2,3,4,5}
	// fmt.Println(array)

	// fmt.Println("slice co phai la Slice khong ? ", reflect.TypeOf(slice).Kind() == reflect.Slice)
	// fmt.Println("array co phai la Array khong ? ", reflect.TypeOf(array).Kind() == reflect.Array)

	// mảng có tổng phần tử là 5
	// [1:4]
	// 1 -> Tính từ vị trí thứ 0 -> n - 1 (n là tổng phần tử của mảng)
	//		=> 0 -> 5 - 1
	//		=> 0 -> 4
	// 4 -> Tính từ vị trí thứ 1 -> n (n là tổng phần tử của mảng)
	//		=> 1 -> 5

	// arr := [5]int{10, 20, 30, 40, 50}
	// fmt.Println("arr co phai la Array khong ? ", reflect.TypeOf(arr).Kind() == reflect.Array)
	// fmt.Println(arr)

	// slice := arr[1:4]
	// fmt.Println(slice)
	// fmt.Println("slice co phai la Slice khong ? ", reflect.TypeOf(slice).Kind() == reflect.Slice)

	// slice := make([]int, 2, 5) // [0, 0]
	// slice[0] = 1		// [1, 0]
	// slice[1] = 2		// [1, 2]
	// slice = append(slice, 3, 4, 5, 6, 7, 8, 9 , 10, 11) // [1, 2, 3, 4, 5, 6]
	// fmt.Println(slice)
	// fmt.Println("Chieu dai cua slice:", len(slice))
	// fmt.Println("Dung luong toi da cua slice:", cap(slice))

	// slice := []int{1,2,3,4,5}
	// fmt.Println(slice)
	// fmt.Println("Chieu dai (length) cua slice:", len(slice))
	// fmt.Println("Dung luong (capacity) toi da cua slice:", cap(slice))

	// apple := []string{"Apple 1", "Apple 2", "Apple 3", "Apple 4", "Apple 5"}
	// fmt.Println(apple[0])
	// fmt.Println(apple[1])
	// fmt.Println(apple[2])
	// fmt.Println(apple[3])
	// fmt.Println(apple[4])
	// fmt.Println(len(apple))

	// for i := 0; i < len(apple); i++ {
	// 	fmt.Println(apple[i])
	// }

	// for key, value := range apple {
	// 	fmt.Printf("apple[%d] = %s \n", key, value)
	// }

	// school := [][]string{
	// 	{"Tuan", "Binh", "Dai"},
	// 	{"Duy", "Hoa", "Huynh"},
	// }

	// fmt.Println(school[1][1])
	// for _, class := range school {
	// 	for _, student := range class {
	// 		fmt.Println(student)
	// 	}
	// }

	// apple := []string{"Apple 1", "Apple 2", "Apple 3", "Apple 4", "Apple 5"}
	// apple = append(apple, "Apple 6", "Apple 7")
	// fmt.Println(apple)

	// apple1 := []string{"Apple 1", "Apple 2", "Apple 3"}
	// apple2 := []string{"Apple 4", "Apple 5", "Apple 6"}
	// apple3 := []string{"Apple 7", "Apple 8", "Apple 9"}
	// apple1 = append(apple1, apple2...)
	// apple1 = append(apple1, apple3...)
	// fmt.Println(apple1)

	// subSlice := slice[1:] // [20, 30, 40, 50]
	// subSlice := slice[:4] // [10, 20, 30, 40]
	// subSlice := slice[1:4] // [20, 30, 40]
	// slice := []int{10, 20, 30, 40, 50}
	// fmt.Println("Slice cha")
	// fmt.Println(slice)
	// fmt.Println("Length slice:", len(slice))
	// fmt.Println("Capacity slice:", cap(slice))

	// fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	// fmt.Println("Slice con")
	// subSlice := slice[2:4]
	// subSlice = append(subSlice, 90, 100)
	// fmt.Println(subSlice)
	// fmt.Println("Length subSlice:", len(subSlice))
	// fmt.Println("Capacity subSlice:", cap(subSlice))

	// TÌM HIỂU HÀM SLICES
	/** Clone: Tạo bản sao của slice **/
	// copied := slices.Clone([]int{1, 2, 3})
	// fmt.Println(copied)

	/** So sánh 2 slice có giống nhau không **/
	// compareSlice := slices.Equal([]int{1, 2, 3}, []int{1, 2})
	// fmt.Println(compareSlice)

	/** Tìm vị trí đầu tiên của phần tử **/
	// findFirstPosition := slices.Index([]int{1, 2, 3, 4, 2}, 2)
	// fmt.Println(findFirstPosition)

	/** Kiểm tra phần tử có trong slice không **/
	// itemExist := slices.Contains([]int{1, 2, 3}, 5)
	// fmt.Println(itemExist)

	/** slices.Insert(s, i, v...) Chèn phần tử vào vị trí i  **/
	// insertValueAnyPostion := slices.Insert([]int{1, 3}, 2, 2)
	// fmt.Println(insertValueAnyPostion)

	/** slices.Delete(s, i, j) Xóa phần tử từ vị trí i đến j-1 **/
	// deleteValueAnyPostion := slices.Delete([]int{1, 2, 3, 4}, 1, 3)
	// fmt.Println(deleteValueAnyPostion)

	/** Đảo ngược slice **/
	// s := []int{1, 2, 3}
	// slices.Reverse(s)
	// fmt.Println(s)

	/** Sắp xếp slice tăng dần **/
	// s := []int{3, 1, 2}
	// slices.Sort(s)
	// fmt.Println(s)

	/** Sắp xếp theo điều kiện tùy chỉnh **/
	// s := []int{3, 1, 2}
	// slices.SortFunc(s, func(a, b int) int {
	// 	return a - b
	// })
	// fmt.Println(s)

	/** Lấy giá trị lớn nhất **/
	// itemMax := slices.Max([]int{1, 3, 2})
	// fmt.Println(itemMax)

	/** Lấy giá trị nhỏ nhất **/
	// itemMin := slices.Min([]int{1, 3, 2})
	// fmt.Println(itemMin)
}
