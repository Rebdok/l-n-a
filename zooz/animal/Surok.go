package animal

import "fmt"

type Surok struct{}

func (c Surok) Move() {
	fmt.Println("Бежиииииит")
}
func (c Surok) Eat() {
	fmt.Println("Кушает не червячка")
}
func (c Surok) Sleep() {
	fmt.Println("Наелся и спит")
}
func (c Surok) Roar() {
	fmt.Println("Чёт тоже орёт")
}
func (c Surok) Backflip() {
	fmt.Println("Сурочья сальтуха")
}
