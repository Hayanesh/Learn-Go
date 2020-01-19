package main

type Person struct {
	Name string
}

func (p *Person) introduce() {
	println("Hi, My name is ", p.Name)
}

func (p *Saiyan) introduce() {
	println("Hi, My Saiyan name is ", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9000,
	}

	goku.introduce()
	goku.Person.introduce()
}
