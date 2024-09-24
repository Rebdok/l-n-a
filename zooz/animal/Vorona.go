package animal

import "fmt"

type Vorona struct{}

func (c Vorona) Move() {
	fmt.Println("Летииииит")
}
func (c Vorona) Eat() {
	fmt.Println("Кушает червячка")
}
func (c Vorona) Sleep() {
	fmt.Println("Наелась и спит")
}
func (c Vorona) Roar() {
	fmt.Println("Чёт орёт")
}
func (c Vorona) Backflip() {
	fmt.Println("Птичья сальтуха")
}
