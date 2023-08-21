package main

// #cgo CFLAGS: -I./c
// #cgo LDFLAGS: -L./c -lperson -Wl,-rpath=./c
// #include "./c/person.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Person struct {
	Age int
	Sex string
}

func (p *Person) IsMale() bool {
	cstr := C.CString(p.Sex)
	defer C.free(unsafe.Pointer(cstr))
	cperson := C.struct_person{
		age: C.int(p.Age),
		sex: cstr,
	}

	result := C.person_is_male(cperson)

	if result == 1 {
		return true
	} else {
		return false
	}
}

func (p *Person) IsAdult() bool {
	cperson := C.struct_person{
		age: C.int(p.Age),
		sex: C.CString(p.Sex),
	}

	result := C.person_is_adult(cperson)

	if result == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Hello world")
	p := Person{
		Age: 20,
		Sex: "female",
	}
	fmt.Printf("Person age=%d sex=%s\n", p.Age, p.Sex)
	fmt.Printf("C calls results\n")
	fmt.Printf("Is Adult? %t\n", p.IsAdult())
	fmt.Printf("Is Male? %t\n", p.IsMale())

}
