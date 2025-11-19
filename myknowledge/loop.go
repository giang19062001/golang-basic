package myknowledge

import "fmt"

func TestLoop() {
	skillList := []string{
		"js", "go", "python", "c#",
	}
	skillTxt := ""
	// for
	for i := 0; i < len(skillList); i++ {
		if i == len(skillList)-1 {
			skillTxt += skillList[i]
		} else {
			skillTxt += skillList[i] + " + "
		}

	}
	fmt.Println("Skills: " + skillTxt)

	// for rang
	for k, skill := range skillList {
		fmt.Printf("skill: %d - %s \n", k, skill)
	}
	//while
	flag := 0
	fmt.Print("while: ")
	for flag != 5 {
		flag++
		fmt.Printf("%d \t", flag)
	}
	//continue
	fmt.Print("\ncontinue: ")
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			continue // bỏ qua ko in các số chia hết cho 2
		}
		fmt.Printf("%d \t", j)

	}

	fmt.Println("\n==============================================")

}
