package main

import "fmt"

func main() {

	for n := 1; n < 100; n++ {
		for i := 0; i < n; i++ {
			if i%5 == 0 {
				fmt.Print("0")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println(n)
	}
}
