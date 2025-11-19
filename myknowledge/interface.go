package myknowledge

import "fmt"

// ** Quy tắc Export trong Go
// Cách viết				Trạng thái				Có thể truy cập từ package khác?
// Name:					Exported (Public)		Có thể dùng ngoài package
// name:					Unexported (Private)	Chỉ dùng trong package hiện tại
// PrintSomething():	 	Public function			Có thể gọi từ file khác
// printSomething(): 		Private function		Chỉ gọi trong package hiện tại

// ! Một struct chỉ tự implement interface khi struct đó có tất cả các method  có trong interface đó
// ! VD: interface có 3 method thì struct củng phải có cả 3 method đó (đúng tên, đúng tham số, đúng kiểu trả về) thì go mới tự

// khái báo interface có hành vi Pay dùng chung cho các struct
type Payment interface {
	Pay(amount float64)
}

// khai báo struct
type CreditCard struct {
	Owner string
}
type PayPal struct {
	Email string
}
type Momo struct {
	Phone string
}

// khai báo method Pay dùng chung, phải cùng tên, cùng param, cùng giá trị trả về, chỉ khác nội dung của method
// như vậy các struct sẽ được go implement Payment interface tự động

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("%s trả %.2f ==> Credit Card\n", c.Owner, amount)
}
func (p PayPal) Pay(amount float64) {
	fmt.Printf("%s trả %.2f ==> PayPal\n", p.Email, amount)
}
func (m Momo) Pay(amount float64) {
	fmt.Printf("%s trả %.2f ==> Momo\n", m.Phone, amount)
}

// khao báo hàm truyền interface (Payment sẽ tự nhận biết là struct nào)
func checkout(p Payment, amount float64) {
	// ! Kiểm tra kiểu thật sự của interface p ( _ là giá trị, ok là bool )
	// ! có thể khai báo biến tạm ngay trong if
	// ! if(biến tạm; logic){}

	// ** cách 1 kiểm tra kiểu của interface dùng if
	if _, ok := p.(CreditCard); ok {
		fmt.Print("CreditCard ==> \t")
	} else if _, ok := p.(PayPal); ok {
		fmt.Print("PayPal ==> \t")
	} else if _, ok := p.(Momo); ok {
		fmt.Print("Momo ==> \t")
	}

	// ** cách 2 kiểm tra kiểu của interface dùng switch
	// switch p.(type) {
	// case CreditCard:
	// 	fmt.Print("CreditCard ==> \t")
	// case PayPal:
	// 	fmt.Print("PayPal ==> \t")
	// case Momo:
	// 	fmt.Print("Momo ==> \t")
	// default:
	// 	fmt.Print("Unknown ==> \t")
	// }

	p.Pay(amount)
}

func TestImplementInterface() {
	cc := CreditCard{"Giang"}
	pp := PayPal{"giang@example.com"}
	mm := Momo{"0909xxx"}

	checkout(cc, 100)
	checkout(pp, 200)
	checkout(mm, 300)
	fmt.Print("==============================================\n")

}
