package main

import (
	"errors"
	"fmt"
)

func main(){
	messages := [3]string{"1", "2", "3"}
	messages[2] = "5"
	
	printMessages(messages)
}
func printMessages(messages [3] string) error{
	if len(messages) == 0{
			return errors.New("Ошибка")
	}
   fmt.Println(messages)

	return nil
}
