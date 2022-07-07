package main

import "fmt"

const (
	catEatsPerMonth = 7.
    dogEatsPerMonth = 10.
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

	var AllAnimals = []Animals{
		Cat{"Kitty", 2},
	    Dog{"Palkan", 10},
		Caw{"Dunysha", 380},
		Dog{"Bob", 14},
		Cat{"Simon", 3},
		Caw{"Toniy", 450},
	} 

	foodForAllAnimals := feedAllAnimals(AllAnimals)
	fmt.Printf("Шоб прогодувати всіх тварин треба %v кг корму \n", foodForAllAnimals)
	
}