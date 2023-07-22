// type NamaStruct struct {
// 	field1 tipeData1
// 	field2 tipeData2
// 	// ...
// 	fieldN tipeDataN
// }

package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Membuat variabel dengan tipe struct Person
    var p1 Person
    p1.Name = "John Doe"
    p1.Age = 30

    fmt.Println("Name:", p1.Name)
    fmt.Println("Age:", p1.Age)
}
package main

import "fmt"

type Rectangle struct {
    width  float64
    height float64
}

// Method untuk menghitung luas dari Rectangle
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Println("Luas:", rect.Area()) // Menggunakan method Area() pada struct Rectangle
}
package main

import "fmt"

type Rectangle struct {
    width  float64
    height float64
}

// Fungsi untuk menginisialisasi Rectangle dengan lebar dan tinggi tertentu
func NewRectangle(width, height float64) Rectangle {
    return Rectangle{width: width, height: height}
}

func main() {
    rect := NewRectangle(10, 5)
    fmt.Println("Lebar:", rect.width)
    fmt.Println("Tinggi:", rect.height)
}
