package main

import "fmt"

func main() {
	var num1, num2, temp int
	fmt.Scanf("%d %d", &num1, &num2)
	temp = num2
	num2 = num1
	num1 = temp

	fmt.Println(num1, num2)
}
