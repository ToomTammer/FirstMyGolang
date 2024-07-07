package toomtam

import "fmt"

func ControlStructure() {
	///IF ELSE
	var score int = 62
 	var grade string

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}

	fmt.Printf("Your grade is %s\n", grade)

	///SWITCH
	var dayOfWeek = 3

	switch dayOfWeek {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Invalid Day")
		}

	///if else สามารถเขียนในรูปย่อแบบนี้ได้
	num1 := 5;
	num2 := 10;

	if sumNum := num1 + num2; sumNum >= 10 {
	fmt.Println("sumNum more than 10")
	}

	//For Loop
	for i := 1; i < 10; i++ {
		fmt.Printf("number: %d\n", i)
	}

	//Do While Loop
	k := 1
	for {
		fmt.Printf("number k: %d\n", k)
		k++
		if k >= 10 {
			break
		}
	}

	j := 1
	//While Loop
	for j < 10 {
		fmt.Printf("number j : %d\n", j)
		j++
	}
}