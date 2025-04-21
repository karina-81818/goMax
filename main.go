package main

import (
	"errors"
	"fmt"
)

func main(){
	message, err := enterGo(10)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	fmt.Println(message)
}

func enterGo(age int) (string, error){

	if age >= 18{
		return "Проходи", nil
	} 
		return "уходи", errors.New("you are too young")
	
}