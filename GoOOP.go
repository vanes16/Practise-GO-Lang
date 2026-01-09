package main

import "fmt"

type Animal interface {
	Info()
	Sound() string
	Rename(name string)
}
type BaseAnimal struct {
	Name string
}

type Cat struct {
	BaseAnimal
}

func NewCat(name string) *Cat {
	return &Cat{BaseAnimal: BaseAnimal{Name: name}}
}
func (b BaseAnimal) Info() {
	fmt.Println("Hello, my name: ", b.Name)
}
func (c Cat) Sound() string {
	return "Meow"
}

func (b *BaseAnimal) Rename(name string) {
	b.Name = name
}

func main() {
	cat := NewCat("Gordon")
	cat.Info()
	fmt.Println(cat.Sound())
	cat.Rename("Alex")
	cat.Info()
}
