package myknowledge

import "fmt"

// defer luôn đảm bảo hàm viết cùng nó luôn gọi khi hàm chứa nó kết thúc, miễn là viết nó ở đầu
// dù là kết thúc do return, panic, hay error,
// trừ khi chương trình chết hẳn bằng os.Exit().

func TestDefer() {

	defer fmt.Println("==============================================")

	defer fmt.Println("Deferred 1")

	fmt.Println("Start")
	defer fmt.Println("Deferred 2")

	fmt.Println("End")

	defer fmt.Println("Deferred 3")

}
