package animal

import "fmt"

type Medoed struct{}

func (c Medoed) Move() {
	fmt.Println("Бежиииииит за кем угодно")
}
func (c Medoed) Eat() {
	fmt.Println("Кушает кого угодно")
}
func (c Medoed) Sleep() {
	fmt.Println("Он никогда не спит")
}
func (c Medoed) Roar() {
	fmt.Println("Орёт на кого угодно")
}
func (c Medoed) Backflip() {
	fmt.Println("Дабл бэкфлип")
}
