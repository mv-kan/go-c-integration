package main

// #cgo CFLAGS: -I${SRCDIR}/c
// #cgo LDFLAGS: -L${SRCDIR}/c -lperson -Wl,-rpath=${SRCDIR}/c
// #include "./c/person.h"
// // or instead of linking shared object you can #include "./c/person.c"
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
	cstr := C.CString(p.Sex)
	defer C.free(unsafe.Pointer(cstr))
	cperson := C.struct_person{
		age: C.int(p.Age),
		sex: cstr,
	}

	result := C.person_is_adult(cperson)

	if result == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	p := Person{
		Age: 20,
		Sex: "female",
	}
	fmt.Printf("Person age=%d sex=%s\n", p.Age, p.Sex)
	fmt.Printf("C calls results\n")
	fmt.Printf("Is Adult? %t\n", p.IsAdult())
	fmt.Printf("Is Male? %t\n", p.IsMale())
	fmt.Println()
	p = Person{
		Age: 10,
		Sex: "male",
	}
	fmt.Printf("Person2 age=%d sex=%s\n", p.Age, p.Sex)
	fmt.Printf("C calls results\n")
	fmt.Printf("Is Adult? %t\n", p.IsAdult())
	fmt.Printf("Is Male? %t\n", p.IsMale())
}
