package main

import "fmt"

func main() {
	fmt.Println("Interface in golang.")
	// interface is set of methods used to define set of actions.

	mike := Student{Human{"Mike",22,"545454"},"MIT",20.23}
	paul := Student{Human{"Paul",19,"232323"},"Harvard",22.00}
	sam := Employee{Human{"Sam",25,"2323211"},"Cloud Factory",23.3}
	tom := Employee{Human{"Tom",32,"67567567"},"HastyAi",232.44}

	// defining interface
	var i Men
	i =mike
	fmt.Println("This is student mike")
	i.sayHi()
	i.Sing("Imagine Dragons")

	// i can store employee
	i= tom
	fmt.Println("This is employee Tom")
	i.sayHi()
	i.Sing("Believer")

	// slice of men

	fmt.Println("Let's use slice of men and see what happens.!!!")
	x := make([]Men,3)
	x[0],x[1],x[2] = paul,sam,mike

	for _,value := range x{
		value.sayHi()
	}

}

type Human struct{
	Name string
	Age int
	Phone string
}

type Student struct{
	Human
	School string
	loan float32
}

type Employee struct{
	Human
	Company string
	money float32
}

// define interfaces
// Interface men is implemented by Human, Employee and Student
type Men interface{
	sayHi()
	Sing(lyrics string)
}

// method
func (h Human) sayHi(){
	fmt.Printf("Hi ! I'm %s and you can call me on %s",h.Name,h.Phone)
}

// method
func (h Human) Sing(lyrics string){
	fmt.Println("alala lala lal lal la lala lalal lala",lyrics)
}


// Employee Overloads method
func (e Employee) sayHi(){
	fmt.Printf("Hi I'm %s and you can call me on %s and I work in %s company",e.Name,e.Phone,e.Company)
}


