package myknowledge

import (
	"fmt"
	"strings"
)

// Interface có  1 hàm Process nhận và xử lý bất kỳ kiểu nào T
type Processor[T any] interface {
	Process(T) T
}

// 1 "hộp" chứa một giá trị kiểu T
type Box[T any] struct {
	Value T
}

// 1 hàm có receiver là "hộp" và param là interface[bất kỳ kiểu nào T]
func (b *Box[T]) Apply(p Processor[T]) {
	b.Value = p.Process(b.Value)
}

// tạo struct để phân biệt [T] + hàm trùng tên trong interface để go tự implement => cho "INT"
type DoubleInt struct{}

func (d DoubleInt) Process(n int) int {
	return n * 2
}

// tạo struct để phân biệt [T] + hàm trùng tên trong interface để go tự implement => cho "STRING"
type UpperString struct{}

func (u UpperString) Process(s string) string {
	return strings.ToUpper(s)
}

func TestGeneric2() {
	defer fmt.Println("==============================================")

	// Box chứa int
	intBox := Box[int]{Value: 10}
	intBox.Apply(DoubleInt{})            // INT
	fmt.Println("intBox:", intBox.Value) // 20

	// Box chứa string
	strBox := Box[string]{Value: "hello"}
	strBox.Apply(UpperString{})          // STRING
	fmt.Println("strBox:", strBox.Value) // "HELLO"
}
