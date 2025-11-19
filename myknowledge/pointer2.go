package myknowledge

import (
	"fmt"
	"sync"
)

type SystemStatus struct {
	mu      sync.Mutex
	clients int
}

func (s SystemStatus) AddClientByValue() {
	fmt.Printf("Địa chỉ của s (bên trong): %p\n", &s)
	s.clients++
}

func (s *SystemStatus) AddClientByPointer() {
	fmt.Printf("Địa chỉ của s (bên trong): %p\n", s)
	s.clients++
}

// ! Khi receiver là giá trị, Go copy struct → vùng nhớ khác nhau.
// * Khi receiver là con trỏ, Go truyền địa chỉ → cùng vùng nhớ

func TestPointer2() {
	defer fmt.Println("==============================================")

	var s SystemStatus
	fmt.Printf("Địa chỉ của s (bên ngoài): %p\n", &s)

	s.AddClientByValue()   // truyền bản sao
	s.AddClientByPointer() // truyền con trỏ
}
