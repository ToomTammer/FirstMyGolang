package main

// main1 for keep code form EP1 beginer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/toomtam/go-example/toomtam"
)

func beginner() {	
	//###Start EP1
	var fullname string = "ToomTam fullname"
	age := 20
	fmt.Printf("Hello %s Yay! age = %d \n",fullname, age)
	id := uuid.New()
	fmt.Printf("UUID: %s\n", id)
	// toomtam.SayTest()
	// toomtam.SayHello()
	
	fullname = "fullname"
	fmt.Printf("Hello %s Yay! \n" ,fullname)

	// toomtam.DataStructure()
	// toomtam.ControlStructure()
	// toomtam.Function()
	// toomtam.Pointer()
	toomtam.CustomError()
	//###End EP1

}