package myknowledge

import "fmt"

func calSalary(salary float64, day int) float64 {
	return salary * float64(day)
}
func getNames() (name string, middleName string, familyName string, age int) {
	// name = "Giang"
	// middleName = "Hoang"
	// familyName = "Ngo"
	// age = 25
	// return name, middleName, familyName, age

	// khai báo 1 dãy hằng số
	const (
		constName       = "Giang"
		constMiddleName = "Hoang"
		constFamilyName = "Ngo"
		constAge        = 25
	)

	return constName, constMiddleName, constFamilyName, constAge
}

func TestPrint() {
	// := Dùng để khai báo một biến MỚI và đồng thời gán giá trị cho nó
	// =  Dùng để gán giá trị cho một biến ĐÃ TỒN TẠI trước đó
	username := "giangngo"
	password := "123456"
	fmt.Printf("account: %v - %v\n", username, password)

	// ** các kiểu Print
	// Println: in xuống hàng sau đó
	// Print: in cùng hàng
	// Printf: in có thể truyền giá trị vào
	// Sprintf: chỉ gộp chuỗi với giá trị truyền vào ko in

	// gọi từng giá trị từ hàm chỉ lấy name, age
	name, _, _, age := getNames()

	// Printf: in trực tiếp ra màn hình
	// fmt.Printf("Name: %s, age: %d\n", name, age)

	// Sprintf: tạo chuỗi mới để lưu với giá trị truyền vào ko in trực tiếp
	info := fmt.Sprintf("Name: %s, age: %d", name, age)
	fmt.Println(info)

	// ** các kiểu Printf
	// %f: float
	// %d: int
	// %v: any
	// %s: string
	// %%: in dấu %

	// khai báo 1 mảng chuỗi
	hobbyList := []string{
		"eat", "sing", "football", "anime",
	}
	// khai báo 1 mảng số
	numFavoriteList := []int{
		7, 12, 3,
	}

	// tính lương theo tháng
	salaryPerDay := 5000
	salaryPerDayFloat := float64(salaryPerDay) // ép kiểu sang float64

	monthDay := 30
	salary := calSalary(salaryPerDayFloat, monthDay)

	// ** các kiểu hay dùng
	// int
	// unit32 : số int ko âm
	// float64
	// string
	// bool

	fmt.Println("total favorites list:", len(numFavoriteList))
	fmt.Println("total hobbies list:", len(hobbyList))

	// fmt.Printf("salary: %f", salary)

	if salary >= 150 {
		fmt.Printf("SALARY: high salary -  %.2f", salary)
	} else if salary >= 100 && salary < 150 {
		fmt.Printf("SALARY: medium salary-  %.2f", salary)
	} else {
		fmt.Printf("SALARY: low salary-  %.2f", salary)
	}

	fmt.Println("\n==============================================")

}
