package myknowledge

import "fmt"

func TestIfSwitch() {
	defer fmt.Println("==============================================")

	day := "Saturday"

	// ---- Cách 1: Dùng if ----
	if day == "Saturday" || day == "Sunday" {
		fmt.Println("[IF] Hôm nay là cuối tuần")
	} else if day == "Monday" || day == "Tuesday" || day == "Wednesday" || day == "Thursday" || day == "Friday" {
		fmt.Println("[IF] Hôm nay là ngày làm việc")
	} else {
		fmt.Println("[IF] Ngày không hợp lệ")
	}

	// ---- Cách 2: Dùng switch ----
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("[SWITCH] Hôm nay là cuối tuần")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("[SWITCH] Hôm nay là ngày làm việc")
	default:
		fmt.Println("[SWITCH] Ngày không hợp lệ ")
	}
}
