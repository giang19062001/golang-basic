package myknowledge

import (
	"fmt"
)

type UEmail struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

// **START: hàm này gọi hàm kia
func formatEmail(index int, el UEmail) string {
	text := fmt.Sprintf("%d => %s@%s.com \n", index, el.Name, el.Email)
	return text
}

func printEmail(emailList []UEmail, formatFunc func(int, UEmail) string) {
	for i, el := range emailList {
		fmt.Printf("%s", formatFunc(i, el))
	}
}

// **END: hàm này gọi hàm kia

// **START: hàm này tạo và trả về hàm mới
func formatPhone(phone string, country string) string {
	result := phone
	if country == "VN" {
		result = "+84" + phone
	} else {
		result = "+96" + phone
	}
	return result
}

func printPhone(emailUser UEmail, formatFunc func(string, string) string) func(string) {
	// Trả về một hàm khác
	return func(phone string) {
		// phone này chính là : emailList[0].Phone của phoneFunc(emailList[0].Phone)
		fmt.Println("Before format:", phone)
		fmt.Println("Formatted:", formatFunc(phone, emailUser.Country))
	}
}

// **END: hàm này tạo và trả về hàm mới

// **START: hàm này gọi đến 1 hàm ẩn danh
func printCountry(f func(string) string, emailUser UEmail) string {
	return f(emailUser.Country)
}

// **END: hàm này gọi đến 1 hàm ẩn danh

func TestAdvancedFunc() {
	emailList := []UEmail{
		{Email: "gmail", Name: "giang", Phone: "334644324", Country: "VN"},
		{Email: "git", Name: "giang", Phone: "334644324", Country: "US"},
		{Email: "company", Name: "giang", Phone: "334644324", Country: "US"},
	}

	// truyền 1 hàm có sẵn làm param cho hàm khác
	printEmail(emailList, formatEmail)

	// truyền 1 hàm trả ra hàm mới, sau đó gọi nó
	phoneFunc := printPhone(emailList[0], formatPhone)
	// như đã khai báo hàn printPhone trả về 1 hàm mới func(string)
	phoneFunc(emailList[0].Phone)

	// truyền 1 hàm ẩn danh vào 1 hàm khác
	resultCountry := printCountry(func(country string) string {
		if country == "VN" {
			return "Viet Name"
		} else {
			return "my"
		}
	}, emailList[0])

	// tạo 1 hàm ẩn danh và gọi ngay lập tức ( gọi ngay lập tức bằng cách thêm () sau hàm ẩn danh hoặc f() sau khi đã tạo hàm ẩn danh )
	func(txt string) {
		fmt.Println("Hàm ẩn danh đã được gọi ngay lập tức, " + txt)
	}("yeah!") // () có thể truyền biến cho hàm ở đây

	fmt.Printf("Country: %s\n", resultCountry)

	fmt.Println("==============================================")
}
