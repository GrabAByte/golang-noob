// GOOD
length := getLength(email)
if length < 1 {
	fmt.Println("Email is invalid")
}

// BETTER
if length := getLength(email); length < 1 {
	fmt.Println("Email is invalid")
}
