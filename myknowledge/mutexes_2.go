package myknowledge

import (
	"encoding/json"
	"fmt"
	"sync"
)

// ! map không được thiết kế thread-safe => panic ngay lập tức khi phát hiện nhiều goroutine đang đọc/ghi map
// * Nếu nhiều goroutine cùng đọc/ghi map cùng lúc → Go runtime sẽ panic để bảo vệ bộ nhớ => phải dùng Mutex

func TestMutexes2() {
	defer fmt.Println("==============================================")

	users := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	addUser := func(name string, stt int) {
		mu.Lock()         // khóa vùng
		defer mu.Unlock() // mở vùng cho goroutin tiếp theo
		users[name] = stt // TODO: biến map này đang bị nhiều goroutin đoc/ghi cùng lúc => dễ panic => cần mutexes
	}

	for i := 0; i < 20; i++ {
		wg.Add(1) // báo thêm có  1 goroutine cần chờ
		go func(i int) {
			defer wg.Done() // báo goroutin này đã xong
			addUser(fmt.Sprintf("user%d", i), i)
		}(i)
	}

	wg.Wait() // Đợi cho tới khi tất cả goroutin xong => thoát wg

	// in kết quả
	data, _ := json.Marshal(users)
	fmt.Println("Danh sách người dùng:", string(data))
}
