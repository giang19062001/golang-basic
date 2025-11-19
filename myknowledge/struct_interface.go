package myknowledge

import "fmt"

type IUserService interface {
	GetName() string
}

type UserService struct {
	Name string
}

func (s *UserService) GetName() string {
	return s.Name
}

// 1. Trả về value
func NewService1() UserService {
	return UserService{Name: "Value"}
}

// 2. Trả về pointer
func NewService2() *UserService {
	return &UserService{Name: "Pointer"}
}

// 3. Trả về interface
func NewService3() IUserService {
	return &UserService{Name: "Interface"}
}

func TestStructInterface() {
	s1 := NewService1()
	s2 := NewService2()
	s3 := NewService3()

	fmt.Println(s1.GetName()) // OK
	fmt.Println(s2.GetName()) // OK
	fmt.Println(s3.GetName()) // OK
}
