// Домашнє завдання по інтерфейсах:
// кожна тварина в залежності від свого типу споживає різку кількість кілограмів їжі на місяць
// собака - їсть 10кг корму на кожні 5кг власної ваги
// кішка - 7кг на кожен кілограм власної ваги
// корова - 25кг на кожен кілограм власної ваги
// на фермі може бути різна кількість собак, кішок та корів
// кажен тип тварини має сам розраховувати для себе вагу корму
// написати динамічну функцію, яка буде виводити в консоль для кожної тварини на фермі її назву, вагу, та скільки їжі треба для того щоб її прогодувати
// ця функція також моє повертати сумму кг корму на всю ферму для подальшого виводу у консоль

package main

import "fmt"

const (
	catEatsPerMonth = 7.
    dogEatsPerMonth = 2. // собака їсть 10 кг корму на 5 кг власної ваги тому : 10 / 5
	cawEatsPerMonth = 25.
)

type Animals interface {
	nicknamer
	mealGetter
	weightGetter
	typer
}
type nicknamer interface {
     getNickname() string
}
type mealGetter interface {
	getQuantityMeal()  float64
}

type weightGetter interface {
	getWeight() float64
}

type typer interface {
	getType() string
}
type Cat struct {
	Nickname string
	Weight float64
}

func (c Cat) getQuantityMeal() float64 {
	return c.Weight * catEatsPerMonth
}

func (c Cat) getNickname() string {
	return c.Nickname
}

func (c Cat) getWeight() float64 {
	return c.Weight
}

func (c Cat) getType() string {
	return "cat"
}

type Caw struct {
	Nickname string
	Weight float64
}

func (cw Caw) getQuantityMeal() float64 {
	return cw.Weight * catEatsPerMonth
}

func (cw Caw) getNickname() string {
	return cw.Nickname
}

func (cw Caw) getWeight() float64 {
	return cw.Weight
}

func (cw Caw) getType() string {
	return "caw"
}
type Dog struct {
	Nickname string
	Weight float64
}

func (d Dog) getQuantityMeal() float64 {
	return d.Weight * dogEatsPerMonth
}

func (d Dog) getNickname() string {
	return d.Nickname
}

func (d Dog) getWeight() float64 {
	return d.Weight
}

func (d Dog) getType() string {
	return "dog"
}

func feedAllAnimals(a []Animals) float64 {
	foodForAll := 0.
	for _,v := range a {
		monthMealCount := v.getQuantityMeal()
		foodForAll += monthMealCount
		fmt.Printf("%v що зветься %v  вагою %v з'їдає за місяць %v кг корму \n", v.getType(), v.getNickname(), v.getWeight(), monthMealCount)
	
	}
	return foodForAll
}


func main() {

	var allAnimals = []Animals{
		Cat{"Kitty", 2},
	    Dog{"Palkan", 10},
		Caw{"Dunysha", 380},
		Dog{"Bob", 14},
		Cat{"Simon", 3},
		Caw{"Toniy", 450},
	} 

	foodForAllAnimals := feedAllAnimals(allAnimals)
	fmt.Printf("Шоб прогодувати всіх тварин треба %v кг корму \n", foodForAllAnimals)
	
}