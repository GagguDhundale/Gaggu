package main

import (
	"fmt"
	"strings"
)

func main() {
	var applicationname = "My Library"
	const totalbooks int = 20
	remainingbooks := totalbooks
	var input string
	var firstname, lastname, email string
	var userbooks int

	books := []string{
		"The Great Gatsby", "1984", "Pride and Prejudice", "The Catcher in the Rye", "Moby-Dick",
		"The Lord of the Rings", "Harry Potter and the Sorcerers Stone", "The Hobbit", "The Alchemist", "The Da Vinci Code",
		"Brave New World", "Fahrenheit 451", "The Kite Runner", "The Book Thief", "Animal Farm",
		"Crime and Punishment", "Dracula", "Frankenstein", "Jane Eyre", "Wuthering Heights",
	}

	greeting(applicationname, totalbooks, remainingbooks)
	fmt.Print("Do you want to proceed? (yes or no): ")
	fmt.Scan(&input)

	// Convert to lowercase to handle different cases
	input = strings.ToLower(input)

	if input == "yes" {
		fmt.Println("Proceeding further ...")
		fmt.Println("Please enter all the following details")

		// Correct the function call and capture returned values
		firstname, lastname, email, userbooks = userinput()


        // will display all the avaible books 
		fmt.Println("\nAvaible Books:")
		for i,book := range books{
			fmt.Printf("%d. %s\n",i+1,book)
		}

		// let user select books 
		selectedBooks := []string{}
		for i := 0 ; i < userbooks ; i++{
			var choice int 
			fmt.Printf("Enter the number of book you want to buy (1-%d): ",len(books))
			fmt.Scan(&choice)
			if choice > 0 && choice <= len(books){
				selectedBooks = append(selectedBooks,books[choice-1])
			}else{
				fmt.Println("Invalid choice, Try again.")
				i--
			}
		}

		// Output the user details
		fmt.Printf("Your first name is: %v\n", firstname)
		fmt.Printf("Your last name is: %v\n", lastname)
		fmt.Printf("Your email address is: %v\n", email)
		fmt.Printf("You hava booked %d books.\n", userbooks)
		for _,book := range selectedBooks{
			fmt.Println("_ "+ book)
		}
	} else {
		fmt.Println("Program ended here ...")
	}

	fmt.Printf("Thank you for using %v for buying books",applicationname)
	fmt.Print("\nAll time you Welcome !")
}

func greeting(applicationname string, totalbooks int, remainingbooks int) {
	fmt.Printf("Welcome to %v site\n", applicationname)
	fmt.Printf("We %v provide you the best possible books on time\n", applicationname)
	fmt.Printf("We have a total of %v books present in our library and currently we have %v books\n", totalbooks, remainingbooks)
}

func userinput() (string, string, string, int) {
	var firstname, lastname, email string
	var userbooks int

	for{

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstname)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastname)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter the number of books you want to buy (1-20) ")
	fmt.Scan(&userbooks)

	if len(firstname) < 2 || len(lastname) < 2{
		fmt.Println("Error: First name and last name must hava atleast 2 letters.")
		continue
	}
	if !strings.Contains(email,"@"){
		fmt.Println("Error: Email must contain '@.")
		continue
	}
	if userbooks<1 || userbooks > 20 {
		fmt.Println("Error. You can only buy book between 1 - 20 range")
		continue
	}
	break

    }
    return firstname, lastname, email, userbooks
}
