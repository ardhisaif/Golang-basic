package src

type DeretBilangan struct {
	Number *int
}

func (deret *DeretBilangan) EvenNumber() []int {
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i%2 == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func (deret *DeretBilangan) OddNumber() []int {
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i%2 != 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func (deret *DeretBilangan) PrimeNumber() []int {
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i == 1 {
			continue
		}

		checkPrime := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				checkPrime++
			}
		}

		if checkPrime == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func (deret *DeretBilangan) FibonacciNumber() []int {
	var arr []int
	for i := 1; i <= *deret.Number; i++ {
		if i == 1 {
			arr = append(arr, 0)
		} else if i == 2 || i == 3 {
			arr = append(arr, 1)
		} else {
			num := arr[len(arr)-1] + arr[len(arr)-2]
			if num < *deret.Number {
				arr = append(arr, num)
			}
		}
	}
	return arr
}