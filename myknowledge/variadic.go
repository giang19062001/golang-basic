package myknowledge

import (
	"encoding/json"
	"fmt"
)

// cách 1 []
func calMoney1(moneys []float64) float64 {
	total := 0.0
	for i := 0; i < len(moneys); i++ {
		total += moneys[i]
	}
	return total
}

// cách 2 ...
func calMoney2(moneys ...float64) float64 {
	total := 0.0
	for i := 0; i < len(moneys); i++ {
		total += moneys[i]
	}
	return total
}

//! ... là đề dàn trải những giá trị phần tử của 1 mảng

type cost struct {
	day   int
	value float64
}

func getCostDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0) // thêm phần tử vào mảng với giá trị 0.0
		}
		costsByDay[cost.day] += cost.value // update lại giá trị  theo day vì giá trị của day hiện tại tương tự như index
	}
	return costsByDay
}

func TestVariadic() {
	moneys := []float64{5.5, 5.1, 5.7}
	total1 := calMoney1(moneys)
	fmt.Printf("Total money 1: %.2f \n", total1)

	total2 := calMoney2(moneys...)
	fmt.Printf("Total money 2: %.2f \n", total2)

	// in mảng
	result := getCostDay([]cost{
		{day: 0, value: 1.5},
		{day: 1, value: 2.5},
		{day: 2, value: 3.5},
	})
	fmt.Printf("getCostDay: %.2f \n", result)

	// In mảng dạng JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Lỗi:", err)
		return
	}

	fmt.Println("getCostDay: " + string(jsonData))

	fmt.Println("==============================================")

}
