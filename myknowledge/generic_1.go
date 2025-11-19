package myknowledge

import "fmt"

// generic có ràng buộc=> cú pháp : [T any]
func splitAnyContraintSlice[T int | string](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// generic => cú pháp : [T any]
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// int
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// string
func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func TestGeneric1() {
	defer fmt.Println("==============================================")

	listInt := []int{1, 2, 3, 4, 5}
	listStr := []string{"a", "b", "c", "d", "e"}

	// int
	intRs1, intRs2 := splitIntSlice(listInt)
	fmt.Printf("splitIntSlice: %v - %v \n", intRs1, intRs2)
	// sintr
	strRs1, strRs2 := splitStringSlice(listStr)
	fmt.Printf("splitStringSlice: %v - %v \n", strRs1, strRs2)
	// generic
	// anyRs1, anyRs2 := splitAnySlice(listInt)
	anyRs1, anyRs2 := splitAnySlice(listStr)
	fmt.Printf("splitAnySlice: %v - %v \n", anyRs1, anyRs2)
	// generic contraint
	// anyRs1, anyRs2 := splitAnySlice(listInt)
	anyContraintRs1, anyContraintRs2 := splitAnyContraintSlice(listInt) // !chỉ có thể truyền int hoặc string như đã khai báo contraint
	fmt.Printf("splitAnySlice: %v - %v \n", anyContraintRs1, anyContraintRs2)
}
