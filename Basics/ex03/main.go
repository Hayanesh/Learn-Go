package main

import "fmt"

type Saiyan struct {
	name  string
	power int
}

func (s *Saiyan) Super() {
	s.power += 10000
}

func main() {
	goku := Saiyan{
		name:  "Goku",
		power: 9000,
	}
	fmt.Println(goku)

	goku = Saiyan{}
	goku.name = "Goku"
	fmt.Println(goku)

	gohan := Saiyan{"Gohan", 8000}
	fmt.Println(gohan)

	gohan.Super()
	fmt.Println(gohan)
}
