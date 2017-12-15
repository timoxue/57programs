package main

import (
	"math"
)

//TipCalculator calculates the tip
type TipCalculator struct {
	billAmount float64
	tip        float64
	tipRate    int
	total      float64
}

func (tipcal *TipCalculator) calculateTip() {
	tipcal.tip = tipcal.billAmount * float64(tipcal.tipRate) / 100
	tipcal.total = tipcal.billAmount + getFixed(tipcal.tip, 2)
}

func getFixed(number float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Trunc((number+0.5/output)*output) / output
}
