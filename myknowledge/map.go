package myknowledge

import (
	"encoding/json"
	"fmt"
)

// `json:"city"` thêm thì sẽ in json ra name thay vì Name
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func nestedMap(users map[string]Person) map[int]map[string]Person {
	nestedUser := make(map[int]map[string]Person)
	// thêm 1 lớp cha là số int
	i := 1
	for userId, userValue := range users {
		nestedUser[i] = map[string]Person{
			userId: userValue,
		}
		i++
	}
	return nestedUser
}
func TestMap() {

	users := map[string]Person{
		"USR001": {Name: "Giang", Age: 25, City: "Hanoi"},
		"USR002": {Name: "Nam", Age: 30, City: "Hue"},
	}

	// thêm 1 key - value mới
	users["USR003"] = Person{Name: "Vũ", Age: 35, City: "Long An"}

	// xóa 1 key - value
	// delete(users, "USR001")

	fmt.Println(users)
	// in map
	for i, usr := range users {
		fmt.Printf("%s - %v \n", i, usr)
	}
	// in map dạng json
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	fmt.Println(string(data)) // In ra JSON một dòng

	//**  user, ok := users["USR001"] => hàm kiểm tra 1 key có tồn tại trong map hay không ; ok là bool
	if user, ok := users["USR001"]; ok {
		data, _ := json.Marshal(user)
		fmt.Println("user find:", string(data), user)
	} else {
		fmt.Println("Không tìm thấy user với key USR001")
	}

	// in nested map dạng json

	dataNested, err := json.Marshal(nestedMap(users))
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	fmt.Println("Nested users: " + string(dataNested))

	fmt.Println("==============================================")

}
