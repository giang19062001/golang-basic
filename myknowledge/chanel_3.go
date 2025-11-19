package myknowledge

import (
	"fmt"
	"time"
)

// Hàm chỉ gửi: nhận channel kiểu `chan<- int`
func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		fmt.Println("sender gửi:", i)
		ch <- i // chỉ được gửi
		time.Sleep(200 * time.Millisecond)
	}
	close(ch) // đóng channel sau khi gửi xong
}

// Hàm chỉ nhận: nhận channel kiểu `<-chan int`
func receiver(ch <-chan int) {
	for {
		val, ok := <-ch // * đọc và kiểm tra trạng thái channel
		if !ok {
			fmt.Println("receiver: channel đã đóng.")
			break
		}
		fmt.Println("receiver nhận:", val)
		fmt.Println("trạng thái channel:", "mở")
	}
	fmt.Println("receiver: đã nhận hết dữ liệu")
}

func TestChanel3() {
	defer fmt.Println("==============================================")

	ch := make(chan int)

	go sender(ch)   // goroutine gửi dữ liệu
	go receiver(ch) // goroutine nhận dữ liệu

	time.Sleep(time.Second) // đợi goroutine chạy xong
}
