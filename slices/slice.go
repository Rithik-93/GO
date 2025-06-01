package main

import "fmt"

func main() {
	fmt.Print(concat([]string{"A", "B"}, []string{"C", "D", "E"}))
}

func concat(s1, s2 []string) []string {
	s1 = append(s1, s2...)
	return s1
}
