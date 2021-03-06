//Question 5
//Author: Kevin Delassus
//Date: 27/09/17
//Write a guessing game where the user must guess a randomly generated number. 
//After every guess the program tells the user whether their number was too high 
//or too low. At the end, the number of tries should be printed. 
//It counts only as one try if they input the same number multiple times 
//consecutively.
//Adapted From: https://adriann.github.io/programming_problems.html


package main

import (
	//Importing the string and the math random package
	"fmt"
	"math/rand"
)

//Array to store guesses
var entry [50]int

func main() {
	
	//Setting Variables
	var guess int
	var ranNum int
	count := 0
	check := true

	//Generating a random number between 1 & 100
	ranNum = rand.Intn(100)
	
	//Displaying Header
	fmt.Println("Please Enter a Single Number between 1 & 100: ")
	fmt.Println("---------------------------------")
	
	//for to check if guess is not equaled to the random number
	for guess != ranNum{

		fmt.Print("-> ")
		//Taking user input
		fmt.Scan(&guess)
		
		//Checking if guess is greater then the random number
		//If it is then message is returned saying you are too high
		if guess > ranNum{

			check = CheckArray(count,guess)

			if check == true{
				fmt.Println("You are to High!")
				
				//fmt.Println(entry[count])
				count++
			}else{fmt.Println("You entered that number already")}
		}
		//Checking if guess is less then the random number
		//If it is then message is returned saying you are too low
		if guess < ranNum{
			
			//Checking if the number has been entered before
			 check = CheckArray(count,guess)

			 //If Number hasnt been entered before then message is printed and count is incremented
			if check == true{
				fmt.Println("You are to Low!")
				
				//fmt.Println(entry[count])
				count++
			}else{fmt.Println("You entered that number already")}
		}
		//if guess equaled to the random number then the program breaks out of the for loop
		if guess == ranNum{
			break
		}
	}
	//Message displaying that you have won the game
	fmt.Println("Correct You Win")
	//Calling function to check how many times user has guessed
	CountTimes(count)
}
//Function to check the array
func CheckArray(count int, guess int) bool{

	entry[count] = guess

	for x := 0; x < count; x++{
		if entry[x] == guess{
			return false
		}
	}
	return true
}
//Function to check how many times user guessed
func CountTimes(c int){
	if c == 0{
		fmt.Println("Woow you guessed correct on the first go!! WELL DONE!!")
	}else{
		c++
		fmt.Println("You Guessed ", c," times!!")
	}
	fmt.Println("---------------------------------")
}