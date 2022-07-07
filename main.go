package main

import "fmt"

const (
	catEatsPerMonth = 7
    dogEatsPerMonth = 10
	cawEatsPerMonth = 25
)

type Animals interface {
	nicknameGetter
	mealGetter
	weightGetter
	typeGetter
}
type nicknameGetter interface {
     getNickname() string
}
type mealGetter interface {
	getQuantityMeal()  int
}

type weightGetter interface {
	getWeight() int
}

type typeGetter interface {
	getType() string
}
type Cat struct {
	Nickname string
	Weight int
}

func (c Cat) getQuantityMeal() int {
	return c.Weight * catEatsPerMonth
}

func (c Cat) getNickname() string {
	return c.Nickname
}

func (c Cat) getWeight() int {
	return c.Weight
}

func (c Cat) getType() string {
	return "cat"
}

type Caw struct {
	Nickname string
	Weight int
}

func (cw Caw) getQuantityMeal() int {
	return cw.Weight * catEatsPerMonth
}

func (cw Caw) getNickname() string {
	return cw.Nickname
}

func (cw Caw) getWeight() int {
	return cw.Weight
}

func (cw Caw) getType() string {
	return "caw"
}
type Dog struct {
	Nickname string
	Weight int
}

func (d Dog) getQuantityMeal() int {
	return d.Weight * dogEatsPerMonth
}

func (d Dog) getNickname() string {
	return d.Nickname
}

func (d Dog) getWeight() int {
	return d.Weight
}

func (d Dog) getType() string {
	return "dog"
}

func main() {
	feedForAllAnimals := 0

	var AllAnimals = []Animals{
		Cat{"Kitty", 2},
	    Dog{"Palkan", 10},
		Caw{"Dunysha", 380},
		Dog{"Bob", 14},
		Cat{"Simon", 3},
		Caw{"Toniy", 450},
	} 

	for _,v := range AllAnimals {
		monthMealCount := v.getQuantityMeal()
		feedForAllAnimals += monthMealCount
		fmt.Printf("%v що зветься %v  вагою %v з'їдає за місяць %v кг корму \n", v.getType(), v.getNickname(), v.getWeight(), monthMealCount)
	}
	// fmt.Printf("\n")
	fmt.Printf("Шоб прогодувати всіх тварин треба %v кг корму \n", feedForAllAnimals)
	
}