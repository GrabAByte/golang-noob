package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	fmt.Println("Authorization: Basic", username+":"+password)
}
