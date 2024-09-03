package main

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// float32 float64
// complex64 complex128

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}
