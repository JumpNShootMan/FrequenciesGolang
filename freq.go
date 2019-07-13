package main

import "fmt"

func freq() {
	//fmt.Println("Hello, world.")
	freq := []int{1, 2, 1, 3, 3, 3, 3, 4, 5, 1, 1, 2, 3, 3, 4, 5}
	freqtotal := []int{0, 0, 0, 0, 0, 0}
	fmt.Println(freq)

	for _, num := range freq {
		freqtotal[num]++
	}

	fmt.Println(freqtotal)

}
