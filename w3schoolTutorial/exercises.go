package main

import "fmt"

func main() {

	//TODO: create a for loop after each execution of functions to add an extra println

	helloWorld()
	varZ()
	varsOfSameType()
	arrayCars()
	goConditions()
	goSwitch()
	goLoop1()
	goLoop2()
	goLoop3()
	goLoop4()
}

// Write a Hello World function
func helloWorld() {
	extraPrintln()
	fmt.Println("Gopher alert 🐀")
	extraPrintln()
}

func extraPrintln() {
	fmt.Println()
}

/*Create a variable called z, assign x + y to it, and display the result.*/
func varZ() {
	x := 5
	y := 10
	z := x + y
	fmt.Println("var z = var x(5) + var y(10) = ", z)
	extraPrintln()
}

// Fill in the missing parts to create three variables of the same type. The three variables should have the following values: 5, 10, and 15.
func varsOfSameType() {
	var x, y, z = 1, 2, 3
	fmt.Println("var x, y, z =", x, y, z)
	extraPrintln()
}

// Create an array of type string called cars.
func arrayCars() {
	cars := [3]string{"Beamer", "Benz", "Bentley"}
	fmt.Println(cars)
	fmt.Println("The second element in the cars array is: ", cars[1])
	fmt.Println("Changing the value from \"Beamer\" to \"Pagani\" in the cars array")
	cars[2] = "Pagani"
	fmt.Println("Last element of cars array: ", cars[2])
	extraPrintln()
}

func goConditions() {
	//TODO: Add in user input for values

	println("Print \"1\" if x is equal to y, print \"2\" if x is greater than y, otherwise print \"3\".")
	x := 50
	y := 50
	if x == y {
		fmt.Println("1")
	} else if x > y {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
	extraPrintln()
}

func goSwitch() {
	println("Use switch to choose between \"Saturday\"(var = 1), \"Sunday\"(var = 2) or \"Weekday\"(var != 1 or 2) by \"day\" variable")
	var day = 4
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	default:
		fmt.Println("Weekday")
	}
	extraPrintln()
}

// Print a list number between 0 & 6
func goLoop1() {
	println("Print i as long as i is less than 6")

	numbers := [6]int{}

	//create array and store numbers instead

	//First initialize the starting position, then where it needs to end, and lastly the operation done to iterate (ex: i++)\
	//Use ";" to seperate the statements in for loop
	for i := 0; i < 6; i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)
	extraPrintln()
}

func goLoop2() {
	fmt.Println("Use a for loop to print \"Yes\" 5 times:")

	for i := 0; i < 4; i++ {
		fmt.Println("Yes")
	}
	extraPrintln()
}

// break & continue
func goLoop3() {
	fmt.Println("Iterate and print through loop from 0 through 10 but stop when it hits 5. Also skip 3")

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}

		fmt.Println(i)
		if i == 5 {
			break
		}
	}
	extraPrintln()
}

// using "range", "idx", "val"
func goLoop4() {
	fmt.Println("Give your top 3 fave ball players of all time using \"range\", \"idx\", \"val\" keyword")

	//create an array of 3 elements
	ballPlayers := [3]string{"Lebron", "Jordan", "Kobe"}

	//create for loop that iterates the array and uses range to output index and value
	for idx, val := range ballPlayers {
		fmt.Printf("%v\t%v\n\n", idx+1, val)
	}
	extraPrintln()
}
