package main

import (
	"fmt"
	"strings"
)

// Type Inference is a feauture to detect variable type by its value.
var conferenceName = "Golang Conf" // := means its a variable and use type inference, it's an e.g. of Syntatic Sugar.
const conferenceTickets int16 = 50 // int16 - datatype for 16bit integer, similar to short in java.
var remainingTickets uint = 50     // uint - datatype for positive whole numbers.
// var bookings [50]string // Array, [fixedSize]datatype - complex datatype for holding multiple ele of same type.
var bookings []string // Slice, dynamic array.
// Package level variables.

func main() {

	greetingUser()

	for remainingTickets > 0 {

		var firstName, sureName, emailAddress = userInfoInput()
		var isValid = validUserInfo(firstName, sureName, emailAddress)
		if !isValid {
			continue
		}
		var userTicketCount = inputTicketAndValidPossiblity()
		var totalTicketCost float64 = 999.99 * float64(userTicketCount)
		remainingTickets = remainingTickets - userTicketCount
		// bookings[0] = userName
		bookings = append(bookings, firstName+" "+sureName)

		fmt.Printf("Hi %s, you have taken %d tickets, and it costs: %.2f\n", firstName, userTicketCount, totalTicketCost)
		fmt.Printf("A comfirmation will be sent to %s, and get the tickets in hand.\n", emailAddress)

		outputShow()
	}

	fmt.Printf("All %s tickets are sold out.\n", conferenceName)
	fmt.Println("Booking System exiting...")
}

//
//
//
//

func greetingUser() {
	fmt.Printf("Welcome to %v Booking System.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and currently %d are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Book as soon as possible, before gets sold out.\n")
}

func userInfoInput() (string, string, string) {
	var firstName string
	fmt.Printf("Enter your first name : ")
	fmt.Scanln(&firstName)

	var sureName string
	fmt.Printf("Enter your sure name : ")
	fmt.Scanln(&sureName)

	var emailAddress string
	fmt.Printf("Enter you email : ")
	fmt.Scanln(&emailAddress)

	return firstName, sureName, emailAddress
}

func validUserInfo(firstName string, sureName string, emailAddress string) bool {
	isVaildName := len(firstName) > 2 && len(sureName) > 2
	isValidEmail := strings.Contains(emailAddress, "@") && strings.Contains(emailAddress, ".") && len(emailAddress) > 6
	if !isVaildName || !isValidEmail {
		if !isVaildName {
			fmt.Println("Name, you entered is too short to true.")
			return false
		}

		if !isValidEmail {
			fmt.Println("Your entered Email is not valid.")
			return false
		}
	}
	return true
}

func inputTicketAndValidPossiblity() uint {
	var userTicketCount uint
	fmt.Printf("How many tickets you need : ")
	fmt.Scanln(&userTicketCount)

	if userTicketCount > remainingTickets {
		fmt.Printf("We have only %v tickets remaining \n", remainingTickets)
		fmt.Printf("How many tickets you need : ")
		fmt.Scanln(&userTicketCount)
		if userTicketCount > remainingTickets {
			fmt.Printf("Your session is terminated...\n")
		}
	}
	return userTicketCount
}

func outputShow() {
	var firstNamesOfBookings []string
	for _, booking := range bookings {
		var names []string = strings.Fields(booking)
		firstNamesOfBookings = append(firstNamesOfBookings, names[0])
	}

	fmt.Printf("System datas : \n Booking list - %v \n Remaining tickets - %d \n", firstNamesOfBookings, remainingTickets)
}
