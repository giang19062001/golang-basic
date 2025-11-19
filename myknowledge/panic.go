package myknowledge

import (
	"fmt"
)

// * panic là một cơ chế để dừng chương trình ngay lập tức khi xảy ra lỗi nghiêm trọng — thường là lỗi mà bạn không thể hoặc không nên xử lý được
// !Chỉ nên dùng panic trong trường hợp lỗi không thể phục hồi
// !Không nên dùng panic để xử lý lỗi thông thường, vì Go khuyến khích xử lý lỗi bằng error thay vì crash chương trình.
// * có thể “chặn” panic để tránh chương trình bị crash bằng recover

// ? Tình huống thật:
// ? Một ứng dụng web cần đọc file cấu hình .env hoặc kết nối DB ngay khi khởi động.
// ? Nếu không được — crash luôn (panic), vì chương trình không thể vận hành trong trạng thái đó.

// Hàm chia có thể gây panic nếu mẫu số = 0
func divide(a, b int) int {
	if b == 0 {
		panic("Không thể chia cho 0") // dừng chương trình lập tức
		// ? log.Fatal("Không thể chia cho 0")
		// ! log.Fatal() y như panic : log và dừng chương trình nhưng ko thể chạy các defer trước cũng như ko thể recover
	}
	return a / b
}

func safeDivide(a, b int) (result int) {
	// Dùng defer + recover để bắt panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Đã bắt panic:", r)
			result = 0 // Gán giá trị mặc định khi lỗi
		}
	}()

	return divide(a, b)
}

func TestPanic() {
	defer fmt.Println("==============================================")

	fmt.Println("Chia 10 / 2 =", safeDivide(10, 2)) // Bình thường
	fmt.Println("Chia 5 / 0 =", safeDivide(5, 0))   // Gây panic nhưng recover lại
	fmt.Println("Chương trình vẫn chạy bình thường sau panic!")
}
