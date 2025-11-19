package myknowledge

import (
	"fmt"
	"sync"
)

// !Nếu nhiều goroutine cùng làm điều này đồng thời, thì ~ goroutine có thể đọc cùng một giá trị cũ.
// !Chúng cùng ghi kết quả mới — nhưng một trong hai kết quả sẽ bị ghi đè mất.
// ? goroutine 1,2,3 đang update counter từ 0 lên 1 song song nhau => goroutine 4 vẫn nhận được counter chỉ là 1
// * => dùng mu.Lock(), mu.Unlock() để khóa vùng truy cập chung => mặc dù giảm tốc độ xử lý nhưng tránh cập nhập trùng
// ! chỉ nên dùng khi nhiều channel cùng cập nhập chung 1 giá trị nào đó

func TestMutexes1() {
	defer fmt.Println("==============================================")

	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1) // báo thêm có  1 goroutine cần chờ
		go func() {
			mu.Lock()         // khóa vùng sau khi có gouroutine nào đó truy cập
			defer mu.Unlock() // mở vùng cho goroutine tiếp theo truy cập
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Kết quả cuối cùng:", counter)
}
