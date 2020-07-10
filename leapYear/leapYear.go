package main

import "fmt"

func main() {
	var year int32
	fmt.Printf("Enter the year")
	fmt.Scanf("%d", &year)

	if year%4 == 0 {
		fmt.Println(year, "is leap year")
	} else {
		fmt.Println(year, " is not leap year")
	}
}
