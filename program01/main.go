package main

import (
	"fmt"
)

func main() {
	fmt.Printf("test")
	totalMoney := &TipCalculator{
		billAmount: 11.25,
		tipRate:    15,
	}
	totalMoney.calculateTip()
	fmt.Printf("Tip is: %.2f\n total Money is: %.2f", totalMoney.tip, totalMoney.total)
}
