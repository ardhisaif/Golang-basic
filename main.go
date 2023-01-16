package main

import (
	"fmt"

	"github.com/ardhisaif/golang_basic/src"
)

func main() {

	//! task 1
	fmt.Println(src.Round(4.37))
	fmt.Println(src.Round(4.32))
	fmt.Println(src.Round(4.324))

	//! task 2
	var q int = 40
	deret := src.DeretBilangan{Number: &q}

	fmt.Println(deret.EvenNumber())
	fmt.Println(deret.OddNumber())
	fmt.Println(deret.FibonacciNumber())
	fmt.Println(deret.PrimeNumber())

	//! task 3
	var s float64 = 6
	sisi := src.Data{Sisi: &s}
	fmt.Println(sisi.Keliling())
	fmt.Println(sisi.Luas())
	fmt.Println(sisi.Volume())

}