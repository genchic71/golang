package main

import "fmt"

// здесь 1 , 2, 3 задачки с 1 уровня
type bankAccount struct {
	balance int
}

func (b *bankAccount) Deposit(amount int) {
	b.balance += amount
}
func (b bankAccount) GetBalance() int {
	return b.balance
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello, I'm", p.Name)
}

func (p *Person) Birthday() {
	p.Age += 1
	fmt.Println(p.Age)
}

//	func (p Person) Birthday() {
//		p.Age += 1
//		fmt.Println(p.Age)
//	}
func main() {
	person := Person{
		Name: "Gena",
		Age:  19,
	}
	balance := bankAccount{
		balance: 20,
	}
	amount := 100
	balance.Deposit(amount)
	fmt.Println(balance.GetBalance())

	person.Greet()
	person.Birthday()
	fmt.Println(person.Age)
}
