package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Printf("Enter range b/w which it needs to find out whether number is even or odd")
	fmt.Scanf("%d %d", &num1, &num2)
	for i := num1; i <= num2; i++ {
		if i%2 == 0 {
			fmt.Printf("Given value %d is even \n", i)
		} else {
			fmt.Printf("Given value %d is odd \n", i)
		}
	}
}
