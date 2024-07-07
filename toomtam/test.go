package toomtam

import "fmt"

func generateTest() {
	fmt.Println("Hello Test!")
}

func SayTest() {
	generateTest()
}

////note
// SayTest() : public function
// sayTest() : private function
// Note สำหรับเรื่องของ package

// 1.go ไม่มี concept class ทุกอย่างจัดการผ่าน function และ package
// 2.ใน 1 folder สามารถมีได้เพียง package เดียวเท่านั้น (หากตั้งชื่อ package ต่างกัน ใน folder เดียวกัน = จะเกิด error ออกมา)
// 3.go มี folder ชื่อ internal ในการจัดการการมองเห็นได้ เช่น หากเรามี library เพิ่มเติมเป็นแบบนี้ เช่นมี folder internal/lopster ที่เป็น package lopster
// ├── go.mod
// ├── go.sum
// ├── main.go
// └── mike
//   ├── internal
//   │   └── lopster
//   │       └── lopster.go
//   ├── mikelopster.go
//   └── test.go