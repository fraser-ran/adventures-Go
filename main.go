package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 10
	var floatNum float64 = 10.2

	fmt.Println("int: ",intNum, "\n", "float: ", floatNum )

	str:= "beans"
	fmt.Println(utf8.RuneCountInString((str)))

	
}