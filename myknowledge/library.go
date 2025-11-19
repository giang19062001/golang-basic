package myknowledge

import (
	"fmt"

	tinytime "github.com/wagslane/go-tinytime"
)

func TestLibrary() {
	defer fmt.Println("==============================================")
	fmt.Println(tinytime.Now())
}
