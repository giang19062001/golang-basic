package myknowledge

import (
	"fmt"
	"sync"
)

type ServerStatus struct {
	mu      sync.Mutex
	running bool
	clients int
}

// ? Nếu bỏ *, chuyện gì xảy ra?
// ! Thì khi gọi s.AddClient(), Go sẽ tạo một bản sao (copy) của ServerStatus, sau đó thực hiện thao tác trên bản sao đó.
// ! Mọi thay đổi (như s.clients++) chỉ xảy ra trên bản copy, không ảnh hưởng đến bản gốc
// *  ====> (s *ServerStatus) thay đổi luôn giá trị gốc trong main

// TODO: (s *ServerStatus) -> Đúng và an toàn -> Mọi goroutine thao tác cùng dữ liệu, mutex chia sẻ
// TODO: (s ServerStatus) -> Không đúng -> Mỗi goroutine có bản sao, mutex bị tách riêng

func (s *ServerStatus) Start() {
	s.running = true
}

// Nếu ko dùng *
// Mỗi goroutine sẽ có bản sao riêng của s, chứa mutex và counter riêng.
// Các goroutine không còn chia sẻ cùng một mutex, nên Lock/Unlock mất tác dụng.
// Cuối cùng s.clients trong biến thật vẫn bằng 0.

func (s *ServerStatus) AddClient() {
	s.mu.Lock()
	s.clients++ // biến được nhiều goroutin truy cập song song
	s.mu.Unlock()
}

func (s *ServerStatus) Info() {
	fmt.Printf("Running: %v | Clients: %d\n", s.running, s.clients)
}

func TestMutexes3() {
	defer fmt.Println("==============================================")

	var s ServerStatus
	s.Start()

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1) // thông báo thêm 1 goroutine cần chờ
		go func() {
			s.AddClient()
			wg.Done() // Thông báo đã xong
		}()
	}

	wg.Wait() // đợi các goroutine done
	s.Info()
}
