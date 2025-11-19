package myknowledge

import (
	"fmt"
	"time"
)

// 1. channel sẽ hoạt động với goroutine mà "goroutine" sinh ra từ : go f()
// 2. channel sẽ hoạt động trong channel của nó nhưng yêu cầu có số lượng buffer : channel := make(chan int, 2)
// ! channel sẽ ko hoạt động trong channel của ko có có số lượng buffer : channel := make(chan int)

func TestChanel2() {
	defer fmt.Println("==============================================")

	// chanel nhận giá trị trừ gotoutine
	ch1 := make(chan string)

	go func() {
		ch1 <- "Hello from goroutine!" // gửi dữ liệu
	}()

	msg := <-ch1 // nhận dữ liệu
	fmt.Println("channel 1:" + msg)

	// chanel ko nhận giá trị từ goroutien , tự set giá trị cho nó nhưng chanel đó có số lượng buffer
	// ! channel chỉ gửi giá trị đi khi giá trị của channel đã đủ số lượng buffer set trước đó

	ch2 := make(chan int, 2) // ? tạo 1 chanel kiểu giá trị int với tối đa chỉ 2 goroutine
	ch2 <- 1
	time.Sleep(time.Second) // ví dụ stop 1s để xem khi nào channel gửi và nhận giá trị
	ch2 <- 2

	fmt.Printf("channel 2: %d \n", <-ch2)
	fmt.Printf("channel 2: %d \n", <-ch2)

	// ! chanel ko nhận giá trị từ goroutien , tự set giá trị cho nó nhưng chanel đó  KO có số lượng buffer
	// ch3 := make(chan int)
	// ch2 <- 1
	// fmt.Printf("channel 2: %d \n", <-ch3) // ! deadlock!
}
