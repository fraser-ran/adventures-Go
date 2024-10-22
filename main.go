package main

import (
	"fmt"
	"unicode/utf8"
	"errors"
)

func main() {
	myString := "resume"
	indexed := myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	
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

func arrsSlMaps(){
	//? integer Array 
	var intArr = [...]int32{1,2,3}
	fmt.Println(intArr)
	//? integer Slice -> think ArrayList 
	intSlice := []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 69)
	//? Map / Dictionary / Hash / HashMap 
	myMap := make(map[string]uint8)
	fmt.Println(myMap)
	otMap := map[string]uint8{"Adam":23, "Sarah": 45}
	//* get value from key 
	fmt.Println(otMap["Adam"])
	//! if no item exists, the map will give the default type -> (unsigned int of bit size 8)= 0 
	fmt.Println(otMap["beans"])
	// ? how we can check if adam is an actual key in the map? 
	ageAdam, ok := otMap["Adam"]
	if ok{
		fmt.Printf("the age is %v \n", ageAdam)
	}else {
		fmt.Println("name not found")
	}
	//? to delete an item from a map
	delete(otMap, "Sarah")
	saAge, oks := otMap["Sarah"]
	if oks{
		fmt.Println(saAge)
	}else {
		fmt.Println("name not found")
	}
	otMap["bon"] = 45
	otMap["Fraser"] = 22
	otMap["69"] = 69

	for name, vals := range otMap{
		fmt.Println("Name:", name, "Values:", vals)
	}

	i := 0
	for {
		fmt.Print(i)
		i++
		if i > 10{
			break
		}
	 }

}