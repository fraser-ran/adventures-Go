package main

import (
	"fmt"
	"unicode/utf8"
	"errors"
)

func main() {
	// var intNum int = 10
	// var floatNum float64 = 10.2

	// fmt.Println("int: ",intNum, "\n", "float: ", floatNum )

	// str:= "beans"
	// // fmt.Println(utf8.RuneCountInString((str)))
	//? integer Array 
	var intArr = [...]int32{1,2,3}
	fmt.Println(intArr)
	//? integer Slice -> think ArrayList 
	intSlice := []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 69)

	// println(y)

	
}

/* @Params: takes a string and prints it out and then 
*/
func printME(printVal string)(string, error){
	var err error 
	if utf8.RuneCountInString(printVal)<4{
		fmt.Println("ERROR")
		err = errors.New("To small string")
		return "error", err
	}
	fmt.Println(printVal)
	
	return printVal, err
}