package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

// package level variables
var conferenceName string = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// slices
var bookings []string

//uint only positive numbers

//array intialize with values
// var bookings = [50]string{"Ayush", "Yadav", "Peter"}

// array intialize
// var bookings [50]string
func main() {

	greetusers()

	//to print their type
	// fmt.Printf("conference tickets is %T, remaining tickets is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//we can also do it like this skipping the inside if else checking
	// for remainingTickets > 0 && len(bookings) < 50 {}
	//infinite loop
	for {

		//taking input
		firstName, lastname, email, userTickets := getUserInput()

		isValidName, isValidemail, isVaidTicket := helper.ValidateUserInput(firstName, lastname, email, userTickets, remainingTickets)

		if isValidName && isValidemail && isVaidTicket {

			bookTicket(firstName, lastname, email, userTickets)

			var firstNames = getFirstNames()

			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				//end loop
				fmt.Printf("our conference is booked\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidemail {
				fmt.Println("email doesn't containn @")
			}
			if !isVaidTicket {
				fmt.Println("number of tickets you entered is invalid")
			}

		}

	}

	// city := "London"

	// switch city {
	// case "new York":
	// 	//executes code for booking new york tickets

	// case "Singapore":
	// 	////executes code for booking Singapore tickets

	// case "London":
	// 	// execso

	// default:
	// 	fmt.Print("no valid city")
	// }

}

func greetusers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("\nGet your Tickets to attend\n\n")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	//if we want to directly print it
	// fmt.Printf("the fist names of bookings are %v\n\n", firstNames)

	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastname string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:")
	fmt.Scan(&lastname)

	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter nmber of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastname, email, userTickets
}

func bookTicket(firstName string, lastname string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	//for array
	// bookings[0] = firstName + " " + lastname

	bookings = append(bookings, firstName+" "+lastname)

	// fmt.Printf("the whole slice is %v\n", bookings)
	// fmt.Printf("the length of slice is %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booing %v tickts. you will recieve a confirmation email at %v\n", firstName, lastname, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
