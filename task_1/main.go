package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		Money      float64 = 23
		PearPrise  float64 = 7
		ApplePrise float64 = 5.99
	)
	var (
		amountPears  int = 8
		amountApples int = 9
	)

	fruitPrise := float64(amountApples)*ApplePrise + float64(amountPears)*PearPrise
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n %v грн.\n", fruitPrise)

	numberPears := math.Round(Money / PearPrise)
	fmt.Printf("2. Скільки груш ми можемо купити?\n %v шт.\n", numberPears)

	numberApples := math.Round(Money / ApplePrise)
	fmt.Printf("3. Скільки яблук ми можемо купити?\n %v шт.\n", numberApples)

	priceFruits := 2*ApplePrise + 2*PearPrise
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?\n", priceFruits <= Money)
}
