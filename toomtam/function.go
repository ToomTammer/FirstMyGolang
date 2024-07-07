package toomtam

import "fmt"

// Define a struct type
type Student struct {
	Name    string
	Weight  int
	Height  int
	Grade   string
}

// Speaker interface
type Speaker interface {
	Speak() string
}
// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}


func Function(){
	sayHelloWith("toom")
	number1 := 3
	number2 := 5
	sumNumber := add(number1, number2)
	fmt.Printf("Function Add :%d\n", sumNumber)

	student := Student{
		Name: "Tam",
		Grade:  "A",
	}

	// Call the FullName method on the Student instance
	info := student.FullName()
	fmt.Println("info of the student:\n", info)

	// interface
	dog := Dog{Name: "Buddy"}

	makeSound(dog)
}

func sayHelloWith(name string) {
	fmt.Printf("Hello %s", name)
}

//return ค่า function
func add(a int, b int) int {
	return a + b
}

///receiver method
// Method with a receiver of type Student
// This method returns the full name of the student
func (s Student) FullName() string {
	return s.Name + " " + s.Grade
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}