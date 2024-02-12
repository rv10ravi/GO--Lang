package main

import (
	"fmt"
	"strings"
)

func main(){

	var str string = "Welcome to GoLang!"

	//1.Length Of the String
	fmt.Println("Length of the String: ", len(str))

	//2.Concatenating Stings
	str1 :="Welcome"
	str2 :="Ravi"
	concat := str1 +" "+ str2

	fmt.Println(concat)

	//3.Conveting to Uppercase
	uppercaseStr := strings.ToUpper(str)
	fmt.Println(uppercaseStr)

	//4.Conveting to Lowercase
	lowerStr := strings.ToLower(str)
	fmt.Println(lowerStr)

	//5.Checking substring existence 
	sentence := "JOhn has a dog names siro"
	keyword:= "siro"

	if strings.Contains(sentence,keyword){
		fmt.Println("Substring", keyword, "found in", sentence)
	}else {
		fmt.Println("Substring", keyword, " not found in", sentence)
	}

	//6.Splitting a string
	csvString:="ravi,john,binny,aryan"
	persons:= strings.Split(csvString,",")
	fmt.Println("Splitted Names:")
	for _, person := range persons{
		fmt.Println("-",person)
	}


}