package main

import "log"

type Weapon struct {
	damage    int
	endurance int
}

type Human struct {
	name   string
	health int
	sword  Weapon
	pos    int
}

func (h *Human) goLeft() {
	log.Println("One step left")
	h.pos--
	log.Printf("Current pos is %d \n", h.pos)
}

func (h *Human) goRight() {
	log.Println("One step right")
	h.pos++
	log.Printf("Current pos is %d \n", h.pos)
}

func (h Human) greeting() {
	log.Printf("Hello! My name is %s \n", h.name)
}

func (h Human) bye() {
	log.Println("Bye!")
}

type Action struct {
	Human
	songName string
}

func (a Action) sing() {
	log.Printf("%s is singing \"%s\"\n", a.name, a.songName)

}

func main() {
	var human = Human{"Victor", 100, Weapon{12, 100}, 0}
	human.greeting()
	human.goLeft()
	human.goRight()
	var action = Action{human, "Riptide"}
	action.sing()
	action.bye()
}
