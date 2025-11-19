package myknowledge

import (
	"fmt"
	"sync"
	"time"
)

// * Dùng Mutex:
//  1. Đọc đồng thời: Không (chỉ "1" goroutine)
//  2. Ghi đồng thời: Không (chỉ "1" goroutine)
//  3. Hiệu năng khi đọc nhiều: thấp
//
// * Dùng RWMutex
//  1. Đọc đồng thời: Có ("NHIỀU" goroutine cùng đọc)
//  2. Ghi đồng thời: Không (chỉ "1" goroutine)
//  3. Hiệu năng khi đọc nhiều: cao

// ! RWMutex: nhiều goroutine có thể đọc cấu hình cùng lúc mà không block nhau, chỉ khi có update mới phải Lock.
var (
	rwData int
	rwMu   sync.RWMutex
)

func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Thông báo đã xong
	for i := 0; i < 3; i++ {
		rwMu.RLock() // khóa vùng
		fmt.Printf("Reader %d đọc dữ liệu: %d\n", id, rwData)
		time.Sleep(100 * time.Millisecond)
		rwMu.RUnlock() // mở vùng
	}
}

func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Thông báo đã xong
	for i := 0; i < 3; i++ {
		rwMu.Lock() // khóa vùng
		rwData += 10
		fmt.Printf(" <-Writer-> %d ghi dữ liệu: %d\n", id, rwData)
		time.Sleep(150 * time.Millisecond)
		rwMu.Unlock() // mở vùng
	}
}

func TestRxMutexes() {
	defer fmt.Println("==============================================")
	var wg sync.WaitGroup

	// Tạo nhiều goroutine đọc
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	// Tạo một goroutine ghi
	wg.Add(1)
	go writer(1, &wg)

	// Chờ tất cả goroutine kết thúc
	wg.Wait()
	fmt.Println("Kết thúc chương trình")
}
