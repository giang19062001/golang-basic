package myknowledge

import "fmt"

type userSalary struct {
	name   string
	salary int
}

func sumSalary(users map[string]userSalary, chanel chan<- int) {
	total := 0
	for _, us := range users {
		total += us.salary
	}
	// gửi giá trị cho chanl
	chanel <- total
	// đóng chanel
	close(chanel)
}
func TestChanel4() {
	defer fmt.Println("==============================================")

	chanel := make(chan int) // thêm số buffer cố định or gắn hàm go f()
	// chanel := make(chan int, 1)

	users := map[string]userSalary{
		"USR01": {name: "giang", salary: 500},
		"USR02": {name: "như", salary: 500},
		"USR03": {name: "nghi", salary: 500},
	}

	// go f()
	go sumSalary(users, chanel)
	fmt.Println("Total salary: ", <-chanel)

}
