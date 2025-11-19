package myknowledge

import (
	"fmt"
	"time"
)

// * Channel là “đường dây liên lạc” giữa các goroutine.
// * Khi gửi: nếu chưa ai nhận, goroutine sẽ dừng lại cho đến khi có người nhận.
// * Khi nhận: nếu chưa ai gửi, goroutine sẽ chờ cho đến khi có giá trị được gửi.
// ! kênh phải luôn được đóng từ phía gửi
type UserEmail struct {
	date string
}

func checkOld(emails []UserEmail) {
	olderChannel := make(chan bool)
	now := time.Now()

	go func() {
		for _, e := range emails {
			// parse string thành time.Time
			emailDate, err := time.Parse("2006-01-02", e.date) // băt buộc format giá trị như này trong go 2006-01-02
			if err != nil {
				fmt.Println("Lỗi parse ngày:", err)
				continue
			}

			// nếu ngày nhỏ hơn hiện tại => là cũ
			if emailDate.Before(now) {
				fmt.Println(e.date + " là email cũ.")
				olderChannel <- true // goroutine gửi giá trị cho channel
			} else {
				fmt.Println(e.date + " là email mới.")
				olderChannel <- false // goroutine gửi giá trị cho channel
			}

		}
		close(olderChannel) // đóng channel sau khi gửi xong

	}()

	// đọc giá trị
	// lặp nhận từng giá trị nhận được từ goroutine của channel
	for isOld := range olderChannel {
		fmt.Println("email is old:", isOld)
	}

}

func TestChanel1() {
	defer fmt.Println("==============================================")

	emails := []UserEmail{
		{date: "2025-08-10"},
		{date: "2025-10-25"},
		{date: "2025-10-26"},
	}
	checkOld(emails)

}
