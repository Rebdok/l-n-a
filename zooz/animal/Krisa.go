package animal

import "fmt"

type Krisa struct{}

func (c Krisa) Move() {
	fmt.Println("Бежиииииит орёт")
}
func (c Krisa) Eat() {
	fmt.Println("Кушает орёт")
}
func (c Krisa) Sleep() {
	fmt.Println("орёт во сне")
}
func (c Krisa) Roar() {
	fmt.Println("Орёт пока орёт")
}
func (c Krisa) Backflip() {
	fmt.Println("не делает, орёт")
}
