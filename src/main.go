package main

/*
#cgo CFLAGS: -I ../include
#cgo LDFLAGS: -L ../lib -lsum
#include "sum.h"
*/
import "C"
import "fmt"
func main() {

	i := C.sum(2, 4)
	fmt.Println("sum: ", i)
}
