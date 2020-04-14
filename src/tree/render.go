package main

import "fmt"

func main() {
	nMax := 55
	for n := 1; n < nMax; n++ {
		r := (nMax - n) / 2
		for i := 1; i < r; i++ {
			fmt.Print(" ")
		}
		for i := 0; i < n; i++ {
			if i%5 == 0 {
				fmt.Fprint("", "0")
			} else {
				fmt.Print("*")
			}
		}
		for i := 1; i < r; i++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
