package animal

import "fmt"

type Medved struct{}

func (c Medved) Move() {
	fmt.Println("Бежиииииит за туристом")
}
func (c Medved) Eat() {
	fmt.Println("Кушает туриста")
}
func (c Medved) Sleep() {
	fmt.Println("В зимней спячке")
}
func (c Medved) Roar() {
	fmt.Println("Орёт в горящей машине")
}
func (c Medved) Backflip() {
	fmt.Println("Медвежья сальтуха")
}
