package main // main package contains multiple go files called modules to be used in goprogram

import "fmt" // module to print
func main() { // entry point at go program
	score := 0

	fmt.Println(" Welcome to quiz game")
	//variables and Datatype//
	//var name string = "Diwakar" // "12234" will be considered as string as it is inside quotes
	// var name int = 456 // any numerical value
	// var name unit =    // unsigned integer means only positive number, no zero and negative numbers
	// var name float64 = 4.42 // any decimal value , 64 means it will use 64 bit memory to store in the computer, it can be posyive or negative
	// var name bool = true // bool is true or false
	//name = 123 // this will throw error as 123 is of type int not string
	//name = "hello"      // value pf the string variable can be change by the vale of the same type
	//country := "Bharat" // here we have not given var keyword and var type as GO automatically guess from the input value type, := tells go tp self assume input type for var
	//fmt.Println(name)
	//fmt.Printf("Hello %v, how are you? Are you from %v ", name, country) // %v is a placeholder for the value of variable put after quotes
	// %d is for decimal notations , %f is for floating point numbers, // https://gobyexample.com/string-formatting
	fmt.Printf("Enter your name: ") // we want to use in same line so didnt use println
	// next libe in print can also be added with \n in the print as below
	//fmt.Print("\nwelcome to the office \n Enter your name: ")
	var name string
	var age int
	fmt.Scan(&name) // & address the memory location of the variable name
	fmt.Printf("Enter your Age: ")
	fmt.Scan(&age)

	fmt.Printf("Hello %v , welcome to the game\n", name)
	//fmt.Println(age >= 10) // other conditional operators are < , > , ==, =!, != .
	if age >= 10 {
		fmt.Print("your are eligible to play")
	} else {
		fmt.Printf("Please come back after %v  years ", 10-age)
		return // return is written to give the output of the function when the function eneds, here it will break out of main function if age is less than 10
		// if age is > 10 then only the program will continue
	}

	fmt.Print("\nwhat is your mobile of choice Asus Rog8 or Samsung S24: ")
	var phone string
	var model string
	fmt.Scan(&phone, &model)
	//fmt.Println(phone, model)

	if phone+" "+model == "Asus Rog8" || phone+" "+model == "asus rog8" {
		fmt.Println("Correct")
		score += 1
		// } else if phone+" "+model == "asus rog8" {
		// 	fmt.Println("Correct")
		// 	score++ // score++ is same as score = score+1 or score += 1
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("How many batteries does %v have: ", phone+" "+model)
	var battery uint
	fmt.Scan(&battery)
	if battery == 2 {
		fmt.Println("Correct")
		score += 1
	} else {
		fmt.Println("Incorrect")
	}
	numofques := 2
	percent := (float64(score) / float64(numofques)) * 100 // int/int= 0, so we need to changet type

	fmt.Printf(" Hi %v, You scored %v out of total %v questions and your percenatge is %v%% .", name, score, numofques, percent)
}
