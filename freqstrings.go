package main

import "fmt"

func freqstrings() {
	//fmt.Println("Hello, world.")
	freq := []string{"jam", "leo", "sol", "elphelt", "bedman", "bedman", "johnny", "elphelt", "jam", "bedman"}
	freqtotal := map[string]int{}
	fmt.Println(freq)

	for _, personaje := range freq {
		freqtotal[personaje]++
	}

	fmt.Println(freqtotal)

}
