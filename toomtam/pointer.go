package toomtam

import "fmt"

// Pointer คือ data type อีก 1 ประเภทที่ใช้สำหรับเก็บ memory address ของตัวแปร ถูกใช้ใช้ใน 2 จุดประสงค์ใหญ่ๆคือ

// Mutability = อนุญาตให้สามารถกลับมาแก้ original data ของตัวแปรตัวนั้นได้ (เพื่อให้เป็นการส่งแบบ pass by reference แทน)
// Efficient large structs = สามารถทำให้ส่งข้อมูลขนาดใหญ่ต่อไปได้ (โดยการส่ง address แทน value ของตัวแปรแทน) เช่น
// การส่ง struct ขนาดใหญ่เข้าไป
// การส่ง config ขนาดใหญ่เข้าไป (เช่น config database, connection ต่างๆ)
type Employee struct {
	Name   string
	Salary int
}

func Pointer(){
	// Declare an integer variable
	x := 10

	// Declare a pointer to an integer and assign it the address of x
	// & address
	// * referance
	var p *int = &x

	// Print the value of x and the value at the pointer p
	fmt.Println("Value of x:", x)  // Output: Value of x: 10
	fmt.Println("Value at p:", *p) // Output: Value at p: 10

	// Modify the value at the pointer p
	*p = 20

	// x is modified since p points to x
	fmt.Println("New value of x:", x) // Output: New value of x: 20

	///###Pass by value และ Pass by reference
	y := 20
	changeValue(y)
	fmt.Println(y) // Output: 20 (x is unchanged)

	///###Example use case ของ pointer
	///###modified ข้อมูลต้นฉบับผ่าน Address
	emp := Employee{Name: "John Doe", Salary: 50000}

	giveRaise(&emp, 5000)
	fmt.Println("After raise:", emp)

	///###Linked List
	var head *ListNode

	prepend(&head, 10)
	prepend(&head, 20)

	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

	//###ดึงพวก config มาใช้
	// Initial configuration
	appConfig := &Config{
		LogLevel: "info",
		Port:     8080,
	}

	fmt.Println("Initial Config:", appConfig)

	// Update configuration
	UpdateConfig(appConfig, "debug", 9000)
	fmt.Println("Updated Config:", appConfig)
}

func changeValue(val int) {
	val = 50
}

func giveRaise(e *Employee, raise int) {
	e.Salary += raise
}

type ListNode struct {
	Value int
	Next  *ListNode
}

// Function to add a node to the front of the list
func prepend(head **ListNode, value int) {
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
}

// Config represents the application configuration
type Config struct {
	LogLevel string
	Port     int
}

// UpdateConfig modifies the provided configuration
func UpdateConfig(c *Config, logLevel string, port int) {
	c.LogLevel = logLevel
	c.Port = port
}