package myknowledge

import (
	"fmt"
	"strings"
)

// ** append() chỉ dùng cho mảng
// ! make(type, length, capacity)
// ! type: kiểu dữ liệu cần tạo (phải là slice, map, hoặc channel)
// ! length: độ dài ban đầu
// ! capacity: (tùy chọn) sức chứa tối đa — chỉ dùng cho slice, Nếu bạn bỏ qua dung lượng, nó sẽ mặc định là độ dài

// array => Độ dài cố định, không thể thêm/bớt phần tử, len = cap => Dữ liệu cố định
func get3Messages() [3]string {
	return [3]string{"i love u", "i love cat", "i love dog"}
}

// slice => số phần tử linh động, Hiệu năng cấp phát bộ nhớ trung bình vì len = cap => Dữ liệu tĩnh nhỏ
func getFullMessages() []string {
	return []string{"i love u", "i love cat", "i love dog"}
}

// slice => số phần tử linh động, Hiệu năng cấp phát bộ nhớ cao vì len luôn nhỏ hơn cap => Dữ liệu lớn hoặc linh hoạt
func getMakeMessages() []string {
	messages := make([]string, 0, 2) // len=0, cap=2
	messages = append(messages, "i love u", "i love cat", "i love dog")
	return messages
}

func TestArray() {
	message := getMakeMessages()
	message = append(message, "I love rabbit") // thêm phần tử vào cuối mảng

	// ! Nếu bạn append thêm phần tử và vượt quá cap, Go sẽ tự động cấp phát mảng mới với dung lượng lớn hơn (thường là gấp đôi)
	// ! xóa, khi xóa thì cap ko giảm đi, nó chỉ luôn tăng
	fmt.Printf("length: %d \n", len(message))   // số phần tử hiện có
	fmt.Printf("capacity: %d \n", cap(message)) // số phần tử tối đa slice có thể chứa trước khi phải cấp phát lại bộ nhớ mới.

	// in toàn mảng từng phần tử
	for i := 0; i < len(message); i++ {
		fmt.Printf("%d - %s\n", i, message[i])
	}
	//  bỏ phần tử đầu VỊ TRÍ 0 -> LẤY CÁC PHẦN TỬ CÒN LẠI TỪ VỊ TRÍ 1 TRỞ ĐI
	fmt.Println("BỎ PHẦN TỬ ĐẦU: ", strings.Join(message[1:], ", "))

	// bỏ phần tử cuối  -> LẤY CÁC PHẦN TỬ CÒN LẠI TRƯỚC DẤU :
	fmt.Println("BỎ PHẦN TỬ CUỐI: " + strings.Join(message[:len(message)-1], ", "))
	// bỏ phần tử cuối  -> LẤY CÁC PHẦN TỬ TỪ VỊ TRÍ ĐẦU ĐẾN VỊ TRÍ 2
	fmt.Println("BỎ PHẦN TỬ CUỐI: " + strings.Join(message[:3], ", "))

	// ! LẤY các phần tử NẰM trong khoảng Ở VỊ TRÍ ĐẦU ĐẾN TRƯỚC VỊ TRÍ CUỐI
	// 0:3 => LẤY TỪ 0,1,2
	// 0:2 => LẤY TỪ 0,1
	// 1:2 => LẤY TỪ 1
	// 1:3 => LẤY TỪ 1,2
	fmt.Println("BỎ PHẦN TỬ ĐẦU VÀ 3: " + strings.Join(message[1:3], ", "))

	//XÓA TẤT CẢ PHẦN TỬ
	fmt.Println("XÓA HẾT :" + strings.Join(message[:0], ", "))

	//HIỆN TẤT CẢ PHẦN TỬ
	fmt.Println("HIỆN HẾT :" + strings.Join(message[0:], ", "))

	// LẤY 0, 1  VÀ ( BỎ 2 ) VÀ TỪ 3 TRỞ ĐI
	message = append(message[:2], message[3:]...)
	fmt.Println("BỎ PHẦN TỬ SỐ 2 RA KHỎI MẢNG :" + strings.Join(message, ", "))

	fmt.Println("==============================================")

	// ? [:0] => XÓA HẾT
	// ? [0:] HOẶC [:] => LẤY HẾT
	// ? [:3] => LẤY TỪ 0 - 2
	// ? [3:] => LẤY TỪ 3 - HẾT
	// ? [3:7] => LẤY TỪ 3 ĐẾN 6

}
