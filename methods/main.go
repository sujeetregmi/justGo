package main

import "fmt"

func main() {
	regmi := User{"hekur", "NYC", "nyc@gmail.com", true, 29}
	fmt.Println(regmi)
	fmt.Printf("Name is: %+v, Age is: %+v \n", regmi.Name, regmi.Age)
	regmi.GetStatus()
	regmi.NewMail()
}

type User struct {
	Name    string
	Address string
	Email   string
	Status  bool
	Age     int
}

func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email ="test@regmi.dev"
	fmt.Println("New email is:",u.Email)
}