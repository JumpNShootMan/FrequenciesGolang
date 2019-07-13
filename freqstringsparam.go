package main

import "fmt"

func freqstringsparam(slice []string) {
	//fmt.Println("Hello, world.")
	//freq := []string{"jam", "leo", "sol", "elphelt", "bedman", "bedman", "johnny", "elphelt", "jam", "bedman"}
	freqtotal := map[string]int{}
	fmt.Println(slice)

	for _, personaje := range slice {
		freqtotal[personaje]++
	}

	fmt.Println(freqtotal)

}
