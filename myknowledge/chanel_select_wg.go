package myknowledge

import (
	"fmt"
	"sync"
	"time"
)

// ** select giống như “switch” nhưng dành cho channel.
// Nó chờ bất kỳ channel nào sẵn sàng, và chạy case tương ứng.
func TestChanelSelect() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(120 * time.Millisecond)
		ch1 <- "Data from channel 1"
		close(ch1)
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Data from channel 2"
		close(ch2)
	}()

	// ! Lệnh select chờ cho đến khi ít nhất 1 case có dữ liệu sẵn.
	// Khi một case được chọn và xử lý xong → select kết thúc (chạy xong 1 lần rồi thoát ra).
	// Vì vậy, nếu bạn muốn nhận nhiều lần (ví dụ cả ch1 và ch2 gửi nhiều giá trị)
	// thì bạn phải đặt select trong vòng lặp (for) để lặp lại hành vi chờ đó.

	// * Lặp cho đến khi cả 2 channel đều đóng
	for ch1 != nil || ch2 != nil {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				fmt.Println("ch1 closed")
				ch1 = nil // tránh panic sau này
				continue
			}
			fmt.Println("Received from ch1:", msg1)

		case msg2, ok := <-ch2:
			if !ok {
				fmt.Println("ch2 closed")
				ch2 = nil // tránh panic sau này
				continue
			}
			fmt.Println("Received from ch2:", msg2)

			// ! Khi KHÔNG NÊN dùng default =>
			// ! Khi cần đảm bảo nhận/gửi đủ dữ liệu
			// * Chỉ dùng  khi channel và dữ liệu của nó không quá quan trọng ko ảnh hưởng luồng chính (ví dụ log)
			// default:
			// 	time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("All senders and receivers completed")

}

// * select: Linh hoạt, xử lý đồng thời nhiều channel và logic phức tạp ->	Cần vòng lặp, code dễ rối khi nhiều kênh
// * WaitGroup: Đơn giản, dễ hiểu khi chỉ cần đợi goroutine hoàn tất ->	Không kiểm soát được từng tín hiệu channel cụ thể
// ! WaitGroup chỉ biết khi goroutine kết thúc, chứ không thể đọc, chờ, hay xử lý từng giá trị mà goroutine gửi qua channel như select
func TestWg() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)

	// báo rằng có 4 goroutine cần chờ: 2 gửi + 2 nhận
	wg.Add(4)

	// Gửi dữ liệu vào ch1
	go func() {
		defer wg.Done() // goroutine này báo hoàn thành
		time.Sleep(120 * time.Millisecond)
		ch1 <- "Data from channel 1"
		close(ch1)
	}()

	// Gửi dữ liệu vào ch2
	go func() {
		defer wg.Done() // goroutine này báo hoàn thành
		time.Sleep(100 * time.Millisecond)
		ch2 <- "Data from channel 2"
		close(ch2)
	}()

	// Nhận dữ liệu từ ch1
	go func() {
		defer wg.Done() // goroutine này báo hoàn thành
		for msg := range ch1 {
			fmt.Println("Received from ch1:", msg)
		}
		fmt.Println("ch1 closed")
	}()

	// Nhận dữ liệu từ ch2
	go func() {
		defer wg.Done() // goroutine này báo hoàn thành
		for msg := range ch2 {
			fmt.Println("Received from ch2:", msg)
		}
		fmt.Println("ch2 closed")
	}()

	wg.Wait() // Đợi cho tới khi tất cả wg.Done() được gọi => thoát wg
	fmt.Println("All senders and receivers completed")
}

func TestChanelSelectWg() {
	defer fmt.Println("==============================================")
	TestChanelSelect()
	fmt.Println("---------------")
	TestWg()

}
