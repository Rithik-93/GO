// package main

// import "fmt"

// func main() {
// 	fmt.Print(concat([]string{"A", "B"}, []string{"C", "D", "E"}))
// }

//	func concat(s1, s2 []string) []string {
//		s1 = append(s1, s2...)
//		return s1
//	}
package main

import (
	"fmt"
)

type User struct {
	Name string
}

func getUser(id int) (*User, error) {
	// Simulating: user not found
	return nil, fmt.Errorf("user with ID %d not found", id)
}

func main() {
	_, err := getUser(42)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User not found.")
	}
}
