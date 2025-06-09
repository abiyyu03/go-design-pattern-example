package decorator

import "fmt"

type Coffee interface {
	GetCost() float64
	GetDescription() string
}

type Espresso struct{}

func (e Espresso) GetCost() float64 {
	return 15000
}

func (e Espresso) GetDescription() string {
	return "Espresso very delicious"
}

type CoffeeDecorator struct {
	Coffee
}

type Milk struct {
	Coffee
}

func (m Milk) GetCost() float64 {
	return m.Coffee.GetCost() + 5000
}
func (m Milk) GetDescription() string {
	return m.Coffee.GetDescription() + ", Milk"
}

type Sugar struct {
	Coffee
}

func (s Sugar) GetCost() float64 {
	return s.Coffee.GetCost() + 2000
}
func (s Sugar) GetDescription() string {
	return s.Coffee.GetDescription() + ", Sugar"
}

type Chocolate struct {
	Coffee
}

func (c Chocolate) GetCost() float64 {
	return c.Coffee.GetCost() + 7000
}

func (c Chocolate) GetDescription() string {
	return c.Coffee.GetDescription() + ", Chocolate"
}

func InitCoffee() {
	var myCoffee Coffee = Espresso{}
	myCoffee = Milk{Coffee: myCoffee}
	myCoffee = Chocolate{Coffee: myCoffee}

	fmt.Println("Pesanan: ", myCoffee.GetDescription())
	fmt.Printf("Total Harga: Rp%.0f\n", myCoffee.GetCost())
}
