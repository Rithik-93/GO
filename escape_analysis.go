// package main

// func main() {
// 	a := NewItem(2, 4)
// 	print(a.x, a.y)
// }

// type Item struct {
// 	x int
// 	y int
// }

// func NewItem(x, y int) *Item {
// 	i := Item{x: x, y: y}
// 	return &i // ← Go moves `i` to the heap automatically
// }

// go build -gcflags="-m" .\escape_analysis.go
