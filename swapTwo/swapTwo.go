package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	fmt.Println(num1, num2)
}
