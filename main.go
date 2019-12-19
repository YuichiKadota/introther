package main

import "fmt"

func main(){
	
	c := add(1,3)
	d := multiprication(2,4)

	fmt.Println(c)
	fmt.Println(d)

}


func add(a,b int) int{

	return a + b
}

func multiprication(a,b int) int{

	return a * b
}