package myknowledge

import "fmt"

// Biến doc là một vùng nhớ sống sót sau khi concater() kết thúc
// → vì hàm con (closure) vẫn đang “giữ” tham chiếu tới nó.
// Mỗi lần gọi wordFunc(), bạn đang tác động lên cùng một biến doc, không phải bản sao mới.
// Đây là lý do doc vẫn nhớ được giá trị trước đó.
// nếu khởi tạo nhiều concater thì các giá trị của mỗi concater sẽ khác nhau

func concater() func(string) string {
	doc := ""
	return func(word string) string {
		doc = doc + word
		return doc
	}
}
func TestClosure() {
	defer fmt.Println("==============================================")
	// gọi concater hàm gán hàm con cho biến wordFunc
	// biến wordFunc đã trở thành 1 hàm
	// ** gọi 1 hàm
	// wordFunc1 := concater()
	// ** gọi 2 hàm cùng lúc
	wordFunc1, wordFunc2 := concater(), concater()
	wordFunc1("I ")
	wordFunc2("I ")

	wordFunc1("LOVE ")
	wordFunc2("Dont ")

	wordFunc1("U ")
	wordFunc2("Love ")

	wordFunc1("SO MUCH")
	wordFunc2("U ")

	fmt.Println(wordFunc1("."))
	fmt.Println(wordFunc2("."))

}
