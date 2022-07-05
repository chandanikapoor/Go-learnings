package main


import (
	"fmt"
	"strings"
)



func main(){
	//1.input the string
	var x string
	compstring := "ian"
	fmt.Printf("Please enter a string:")
	_  , err := fmt.Scan(&x)
	if err == nil{
	
		//2.we will compare it with if condition
			if strings.Contains(x, compstring){
			fmt.Printf("Found!")

		}else{
			fmt.Printf("Not Found!")
		}
		
	}else{
		fmt.Printf("Invalid input")
	}

	
}