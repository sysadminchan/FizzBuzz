package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Bizz")
		} else if i&3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Printf("This number is %d\n", i)
		}
	}
}
