package main

import "fmt"

func main() {

	cards := newDeckFromFile("my_cards")

	cards.print()

	fmt.Println("After shuffle ---")

	cards.shuffle()
	cards.print()
}
