package section8

import "fmt"

type Dog1 struct {
	name   string
	weight int
}

type Cat1 struct {
	name   string
	weight int
}

func (d Dog1) bite() {
	fmt.Println(d.name, " - Dog bite!!!")
}

func (d Dog1) sounds() {
	fmt.Println(d.name, " - Dog barks!!!")
}

func (d Dog1) dogFeature() {
	fmt.Println(d.name, " - dogs are loyal")
}

func (d Cat1) bite() {
	fmt.Println(d.name, " - Cat crawls!!!")
}

func (d Cat1) sounds() {
	fmt.Println(d.name, " - Cat cries!!!")
}

func (d Cat1) catFeature() {
	fmt.Println(d.name, " - cats are independent")
}

type Behavior1 interface {
	bite()
	sounds()
}

func act1(animal Behavior1) {
	animal.bite()
	animal.sounds()
}

func Interface3() {

	dog := Dog1{"poll", 10}
	cat := Cat1{"nabi", 3}

	animals := []Behavior1{dog, cat}

	for _, val := range animals {
		act1(val)
	}

}
