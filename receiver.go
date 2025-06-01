package main

import "fmt"

type Item struct {
	x, y int
}

func (i Item) Move(a , b int) {
	
}

// func (i *Item) Move(x, y int) {
// 	i.x = x
// 	i.y = y
}

func main() {
	i3 := Item{x: 2, y: 3}
	i3.Move(100, 200)
	fmt.Println(i3) // Prints: {100 200}
}
