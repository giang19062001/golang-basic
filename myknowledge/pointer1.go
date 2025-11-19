package myknowledge

import "fmt"

func TestPointer1() {
	defer fmt.Println("==============================================")
	a := 10 // giá trị ban đầu
	b := &a // b là con trỏ trỏ đến a (lưu địa chỉ của a)

	fmt.Println("Giá trị a:", a)
	fmt.Println("Địa chỉ của a:", &a)
	fmt.Println("Giá trị con trỏ b (địa chỉ của a):", b)
	fmt.Println("Giá trị mà b trỏ tới (*b):", *b)

	*b = 20 // thay đổi giá trị tại vùng nhớ mà b trỏ tới
	fmt.Println("Sau khi đổi *b = 20, a =", a)
}
