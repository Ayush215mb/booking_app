package helper

import "strings"

func ValidateUserInput(firstName string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2

	isValidemail := strings.Contains(email, "@")

	isVaidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidemail, isVaidTicket

}
