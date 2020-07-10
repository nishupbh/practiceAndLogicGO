package main

import "fmt"

func main(){
	var num int 
	fmt.Printf("Enter the number")
	fmt.Scanf("%d",&num)

	if num % 2 == 0{
		fmt.Printf("Number %d is even number\n", num)
	}else {
		fmt.Printf("Number %d is odd number\n",num)
	}
}