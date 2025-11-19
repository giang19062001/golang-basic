package myknowledge

import (
	"fmt"
	"time"
)

// * Goroutine là một luồng nhẹ (lightweight thread) trong Go.
// * Nó cho phép bạn chạy các hàm song song (concurrently)
// * go func() {
// code chạy song song
// }()

// ! Một chương trình Go tự động thoát khi main() kết thúc, dù còn goroutine chưa chạy xong.
// ! → Vì vậy bạn phải chờ nếu cần chúng hoàn tất (bằng WaitGroup, channel, hoặc Sleep).

type email struct {
	username string
}

func sendEmail(message string) {
	// 1. hàm này vẫn được gọi như thứ tự nhưng ko chờ hàm này trả kết quả xong mà goroutine nhảy xuống thưc thi hàm tiếp theo
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: %s \n", message)
	}()
	// 2. goroutine thực thi hàm này tiếp theo
	fmt.Printf("Email sent: %s \n", message)
	// 3. goroutine thấy khi hàm "received" đã thưc thi xong hàm trên cùng tự xuất ra kết quả

	// * Email sent: giang
	// * Email received: giang
}

func TestGoroutine1() {
	defer fmt.Println("==============================================")
	emails := []email{}
	emails = append(emails, email{username: "giang"})
	emails = append(emails, email{username: "như"})
	emails = append(emails, email{username: "nghi"})

	// in email
	for _, em := range emails {
		sendEmail(em.username)
	}

	time.Sleep(time.Second) // Chờ goroutine chạy xong
}
