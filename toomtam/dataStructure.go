package toomtam

import "fmt"

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Another struct type used in Person
type Address struct {
	Street  string
	City    string
	ZipCode int
}



func DataStructure() {
	//##Array 
	fmt.Println("##Array")
	var myArray [3]int // An array of 3 integers
	myArray[0] = 10    // Assign values
	myArray[1] = 20
	myArray[2] = 30
	fmt.Println(myArray) // Output: [10 20 30]
	// Reassigning the elements of the array
	myArray[0] = 100
	myArray[1] = 200
	myArray[2] = 300

	fmt.Println(myArray) // Output: [100 200 300]

	// Looping through the array
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	//##Slice
	fmt.Println("##Slice")
	mySlice := []int{10, 20, 30, 40, 50} // A slice of integers

	fmt.Println(mySlice)          // Output: [10 20 30 40 50]
	fmt.Println(len(mySlice))     // Length of the slice: 5
	fmt.Println(cap(mySlice))     // Capacity of the slice: 5
  
	// Slicing a slice
	subSlice := mySlice[1:3]      // Slice from index 1 to 2
	fmt.Println(subSlice)	 // Output: [20 30]

	var mySlice2 []int // Declared but not initialized

	// Appending data to the slice
	mySlice2 = append(mySlice2, 10)
	mySlice2 = append(mySlice2, 20, 30)

	fmt.Println(mySlice2) // Output: [10 20 30]

	//Array สามารถ convert มาเป็น slice ได้
	var myArray2 [3]int // An array of 3 integers
	myArray2[0] = 10    // Assign values
	myArray2[1] = 20
	myArray2[2] = 30

	// Converting array to slice
	mySlice3 := myArray2[:]

	// Resizing slice by appending new elements
	mySlice3 = append(mySlice3, 40, 50)

	fmt.Println(mySlice3) // Output will show a slice with 5 elements: [10 20 30 40 50]

	//##Map
	fmt.Println("##Map")
	myMap := make(map[string]int)

	// Add key-value pairs to the map
	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])

	// Update the value for a key
	myMap["banana"] = 12

	// Delete a key-value pair
	delete(myMap, "orange")

	// Iterate over the map
	for key, value := range myMap {
	fmt.Printf("%s -> %d\n", key, value)
	}

	// Checking if a key exists
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value:", val)
	} else {
		fmt.Println("Pear not found inmap")
	}

	//##Struct
	fmt.Println("##Struct")
	var student1 Student
	student1.Name = "Mikelopster"
	student1.Weight = 60
	student1.Height = 180
	student1.Grade = "F"

	// Print struct values
	fmt.Println(student1)
}

func StructOther(){
	//ARRAY
	// Create an array of Student structs
	var students [3]Student

	// Initialize the first student
	students[0] = Student{
		Name:   "Mikelopster",
		Weight: 60,
		Height: 180,
		Grade:  "F",
	}

	// Initialize the second student
	students[1] = Student{
		Name:   "Alice",
		Weight: 55,
		Height: 165,
		Grade:  "A",
	}

	// Initialize the third student
	students[2] = Student{
		Name:   "Bob",
		Weight: 68,
		Height: 175,
		Grade:  "B",
	}

	// Print array of structs
	fmt.Println(students)

	//MAP
	// Create a map with string keys and Student struct values
	students2 := make(map[string]Student)

	// Add Student structs to the map
	students2["st01"] = Student{Name: "Mikelopster", Weight: 60, Height: 180, Grade: "F"}
	students2["st02"] = Student{Name: "Alice", Weight: 55, Height: 165, Grade: "A"}
	students2["st03"] = Student{Name: "Bob", Weight: 68, Height: 175, Grade: "B"}

	// Print the map
	fmt.Println("Students2 Map:", students2)

	// Access and print a specific student by key
	fmt.Println("Student st01:", students2["st01"])


	//STRUCT
	// Create an instance of the Person struct
	var person Person
	person.Name = "Alice"
	person.Age = 30
	person.Address = Address{
		Street:  "123 Main St",
		City:    "Gotham",
		ZipCode: 12345,
	}
	
	
	// Alternative way to initialize a struct
	bob := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			Street:  "456 Elm St",
			City:    "Metropolis",
			ZipCode: 67890,
		},
	}

	// Print struct values
	fmt.Println(person)
	fmt.Println(bob)
}