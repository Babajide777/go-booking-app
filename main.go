package main

import "fmt"

func main() {

	const conferenceName string = "Go conference"
	conferenceTickets := 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, but only %v are availble\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the amount of tickets you want to purchase:")
	fmt.Scan(&userTicket)

	fmt.Println(&firstName)
	fmt.Println(&lastName)
	fmt.Println(&email)
	fmt.Println(&userTicket)

}
