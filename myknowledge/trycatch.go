package myknowledge

import (
	"fmt"
)

// ! vì error trong go là interface có method Error nên các struct nào có method là Error() thì go tự hiểu
// ! là đang implement error interface đó để tạo custom error
// type error interface {
// 	Error() string
// }

type uError struct {
	NotFoundName string
}
type user struct {
	name string
	age  int
}

func (e uError) Error() string {
	return fmt.Sprintf("User '%s' not found", e.NotFoundName)
}

func getUser(name string) (user, error) {

	// ** cách 1
	userList := []user{
		{name: "Giang", age: 25},
		{name: "Nam", age: 30},
		{name: "Hoa", age: 22},
	}

	// ** cách 2
	// userList := []user{}
	// userList = append(userList, user{name: "Giang", age: 25})
	// userList = append(userList, user{name: "Nam", age: 30})
	// userList = append(userList, user{name: "Hoa", age: 22})

	// Duyệt tìm user theo name
	for _, u := range userList {
		if u.name == name {
			return u, nil
		}
	}

	// ! Dùng error mặc định ===> Errorf  ghi lỗi truyền vào error ko in trực tiếp
	// return user{}, fmt.Errorf("user not found - 1")
	// ! Dùng error mặc định ===> errors.New  từ import "errors"
	// return user{}, errors.New("user not found - 2")
	// ! Dùng error custom => tự động gọi hàm Error() của struct đó
	return user{}, uError{NotFoundName: name}
}

func TestTryCatch() {

	u, err := getUser("Giangs")

	// Go xử lý lỗi qua giá trị trả về (error)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("==============================================")
		return
	}
	fmt.Printf("User: %s - Age: %d\n", u.name, u.age)
	fmt.Println("==============================================")
}
