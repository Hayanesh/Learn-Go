# GoLang Learning Notes

## Introduction
It is a statically typed, compiler based language with garbage collection.

### Running Go code

``go run main.go``

Compiles and executes the code

``go run --work main.go``

Shows the location fo the compiled code

``go build main.go``

To explictly compile the code and generate the executable

```
package main

func main(){
    //Entry point for any Go program
}
```

### Imports

It is used to declare the packages that are used by the code in the file

```
import (
    "fmt"
    "os"
)
```

> Go is strict about importing packages. It will not compile if you import a package but don't use it

### Variables and Declarations

```
// Declared variable power of type int and initialize it to 9000
var power int  = 9000

// Short variable declaration
power := 9000

// COMPILER ERROR:
// no new variables on left side of :=
power := 9001

// Assign multiple variables 
name, power := "test" , 9000
=> Here power is re-declared and intialized to 9000. But the Go compiler will not throw any errors, as long as atleast one variable is newly declared. Any redeclaration will be igored. 

println(power) 
=> will return 9001

```

> Only for declaration ``:=`` shall be used and all subsequent assignment operator ``=`` should be used
> Go will throw compilation error if any unused variable is present in the code

### Function Declarations

Functions can return multiple values.

```
// Function with no return value
func log(message String){

}

// Function with one return value
func add(a int, b int) int{

}

// Function with two return value
function power(name string) (int, bool){

}
```

At times when we care about only one of the returned values, we can use the blank identifier ``_,`` to keep the other one unassigned.

## Structures

Go isn't an Object-oriented language, so it doesn't have objects or inheritance or neither support polymorphism or overloading. But rather support composition.

**_Addition references_**

[Inheritance vs Composition](https://www.javaworld.com/article/3409071/java-challenger-7-debugging-java-inheritance.html)

### Declaration and initialization
```
// Declaration
type Saiyan struct {
    name string
    power int
}

// Initialization **
goku := Saiyan{
    name = "Goku",
    power = 9000,
}

(or)

goku := Saiyan{"Goku",9000}

// Initializing to default values
goku = Saiyan{}
goku.name = "Goku"

```

> ** Structure parameters should always have a tailing comma during explicit paramter intialization


### Functions on Structures

```
// Pointer variable intialization
goku :=  &Saiyan{"Goku", 9000}


func Super (s *Saiyan){
    s.power += 10000
}
Super(goku)

(or)

func (s *Saiyan) Super() {
    s.power += 10000
}
goku.Super()

```
> In the above code, we say that the type *Saiyan in the **receiver** of the Super method.


### Constructors

Structures don't have constructors. Instead we can have a function that returns the new instance of the desired type

```
func NewSaiyan(name string, power int) *Saiyan {
    return &Saiyan{
        name: name,
        power: power,
    }
}
```

### New

Despite the lack of constructores, Go does have a built-in new function that is used ot allocate the memory for the memory required by the type

```
goku := new(Saiyan)
// Same as
goku := &Saiyan{}

```

### Fields of a Structure

Fields can be of any type - including other stucture and types like array, map, interface and functions.

```
type Saiyan struct {
    Name string,
    Power int,
    Father *Saiyan
}

gohan := &Saiyan{
    Name: "Gohan",
    Power: 1000,
    Father: &Saiyan{
        Name: "Goku",
        Power: 9000,
        Father: nil,
    },
}

```

### Composition

```
type Person struct {
	Name string
}

func (p *Person) introduce() {
	println("Hi, My name is ", p.Name)
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
    // Same as
    goku.Person.introduce()
}

```

### Overloading
Go doesn't support overloading. However, because **implicit composition** is just a compiler trick, we can "overwrite" the functions of the composed type.

```

type Person struct {
	Name string
}

func (p *Person) introduce() {
	println("Hi, My name is ", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func (p *Saiyan) introduce() {
	println("Hi, My Saiyan name is ", p.Name)
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9000,
	}

	goku.introduce()
    // prints Hi, My Saiyan name is Goku

    goku.Person.introduce()
    // prints Hi, My name is Goku

}
```

## Maps, Arrays and Slices