package main

import (
	"bufio"
	"os"
	"papapap/src/hello-go/zooz/animal"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a animal.Animal

	switch input {

	case "Vorona":
		a = animal.Vorona{}

	case "Surok":
		a = animal.Surok{}

	case "Medved":
		a = animal.Medved{}

	case "Medoed":
		a = animal.Medoed{}

	case "Krisa":
		a = animal.Krisa{}

	}

	a.Move()
	a.Eat()
	a.Sleep()
	a.Roar()
	a.Backflip()

}
