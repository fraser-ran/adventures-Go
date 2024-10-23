package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)
// ? Example of a struct
type gasEngine struct{
	kpl uint8 //km per liter
	liters uint8
	owner 
	//! here we put the owner type so it makes an owner struct that we can directly call stuff from, treating the 
	// ! owner struct that is part of the gasEngine as just extra parts of the gasEngine struct 
}
// ? Here we tack on a method onto the gasEngine struct so we can get some object functionanlity of an object 
func (g gasEngine) fuelRange() uint8{
	return g.kpl*g.liters
}

type owner struct{
	name string
}

func main() {
	structs()
}
func structs(){
	var eEngine gasEngine
	fmt.Println(eEngine.kpl, eEngine.liters, eEngine.name)//! is an empty struct, prints to default values ->(0) 
	// ? when instantiating a struct, the order the variables were created is the default order they will be assigned 
	fullEngine := gasEngine{25, 10, owner{"Fraser"}} 
	fmt.Println(fullEngine.kpl, fullEngine.liters, fullEngine.name)
	fmt.Printf("%v's car has a range of %d km", fullEngine.name, fullEngine.fuelRange())
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
/*String practice*/
func myStrings(){
	strSlice := []string{"s","u","b","o","p","t","i","m","a","l"}
	catStr := ""
	//! very ineffcient method of concatenating a string (creates a new string each time it loops)
	for i:= range strSlice{
		catStr += strSlice[i]
		fmt.Println(catStr)
	}
	opSlice := []string{"o","p","t","i","m","a","l"}
	var strBuilder strings.Builder
	//* optimal method
	for i := range opSlice{
		strBuilder.WriteString(opSlice[i])
	}
	goodCatStr := strBuilder.String()
	fmt.Println(goodCatStr)
}

/*Loop practice*/
func loops(){
	// ? while loop
	i := 0
	fmt.Println("while loop: ")
	for {
		fmt.Print(i)
		i++
		if i > 10{
			break
		}
	 }
	 // ? FOR LOOP 
	 fmt.Println("\nfor loop: ")
	 for i := 0 ; i<4; i++ { 
		fmt.Print(i)
	}

	// ? INCREMENT THROUGH A LIST, SLICE OR STRING LOOP
	fmt.Println("incrememnt through a loop and DO NOT show values")
	myString := "resume"
	indexed := myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v:= range myString{
		fmt.Println(i, v)
	}
	fmt.Println("\n incrememnt through a loop and show values")
	sString := "My String"
	for i, char := range sString{
		fmt.Printf("Index: %d, character: %c \n" ,i, char )
		//! we need to display the char as %c or it will default to its uInt utf-8 val 
	}
}
/*arrays, slices and maps*/
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