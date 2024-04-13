package main

import (
	"fmt"

	"money"
)

func main() {
	rubles := money.FromKopecks(120)

	if rubles == money.RUB(1.2) {
		fmt.Printf("%v rubles is %v kopecks", rubles, rubles.Kopecks())
	}
}

// layout: 1.2 rubles is 120 kopecks
