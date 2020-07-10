package main

import "fmt"

func main() {
	var a, b, c, largest int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b {
		largest = a
		if largest > c {
			fmt.Println("largest number is ", largest)
		} else {
			largest = c
			fmt.Println("largest number is ", largest)
		}
	} else {
		largest = b
		if largest > c {
			fmt.Println("largest number is ", largest)
		} else {
			largest = c
			fmt.Println("largest number is ", largest)
		}
	}

}
