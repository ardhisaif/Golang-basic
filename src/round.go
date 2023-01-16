package src

import (
	"fmt"
	"math"
)

func Round(num float64) string {
	num = math.Round(num*10) / 10
	s := fmt.Sprintf("%.2f", num)
	return s
}
