package main

import "fmt"

func main() {
	p := Player{Name: "Rithik"}

	fmt.Println(p.FoundKey(Jade))    // ✅ adds jade
	fmt.Println(p.FoundKey(Smith))   // ✅ does nothing (already there)
	fmt.Println(p.FoundKey(Key(99))) // ❌ invalid key

	fmt.Println(p.Keys)
}

type Key int

const (
	Jade Key = iota + 1
	Cooper
	Smith
)

type Player struct {
	Name string
	Keys []Key
}

func isValidKey(k Key) bool {
	switch k {
	case Cooper, Jade, Smith:
		return true
	}
	return false
}

func (p *Player) FoundKey(k Key) error {
	if !isValidKey(k) {
		return fmt.Errorf("invalid key: %v", k)
	}

	for _, key := range p.Keys {
		if key == k {
			return nil
		}
	}

	p.Keys = append(p.Keys, k)

	return nil
}
