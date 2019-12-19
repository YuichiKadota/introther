package main

import "fmt"

func main(){
	
	c := add(1,3)

	fmt.Println(c)

}


func add(a,b int) int{

	return a + b
}