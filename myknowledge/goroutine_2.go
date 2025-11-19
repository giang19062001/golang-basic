package myknowledge

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	fmt.Println(msg)
	time.Sleep(250 * time.Millisecond)
}

func TestGoroutine2() {
	defer fmt.Println("==============================================")

	// ! do là chạy song song nên đôi lúc sẽ ra Goroutine 2 đôi lúc ra Goroutine 1 trước,
	// ! tuy nhiên Main routine luôn chạy do ko phải chạy song song
	go printMessage("Goroutine 1") // tạo một goroutine mới -> chạy song song hàm này -> và tiếp tục chạy dòng tiếp theo ngay lập tức.
	go printMessage("Goroutine 2") // tạo một goroutine mới -> chạy song song hàm này -> và tiếp tục chạy dòng tiếp theo ngay lập tức.
	printMessage("Main routine 1")
	printMessage("Main routine 2")
	printMessage("Main routine 3")
	printMessage("Main routine 4")
	printMessage("Main routine 5")

	time.Sleep(time.Second) // Chờ goroutine chạy xong
}
