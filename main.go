package main

import (
	"errors"
	"fmt"
)

func main(){
	 //messages := make([] string, 10000) 
	
	
	//printMessages(messages)
	//fmt.Println(len(messages))
	//fmt.Println(cap(messages))
	//fmt.Println(messages)
	//messages = append(messages, "10001")
	//fmt.Println(len(messages))
	//fmt.Println(cap(messages))
	matrix := make([][] int, 10)
	counter := 0
	for x := 0; x < 10; x++{
		matrix[x] = make([] int, 10)
		for y := 0; y < 10; y++{
			counter++
			
			matrix[x][y] = counter
		}
		fmt.Println(matrix[x])
	}

}
func printMessages(messages [3] string) error{
	if len(messages) == 0{
			return errors.New("Ошибка")
	}


   

	return nil
}
