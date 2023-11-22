package main

import "fmt"

func main() {
	var conferenceName string = "Go Conferenece"
	const confrenceTicket int = 50
	var remainingTicket uint = 50

	remainingTicket = 100

	fmt.Printf("conferenceTickets is %T remainingTickets is %T conferenceName %T\n", confrenceTicket, remainingTicket, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confrenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	//ask user for their name
	userName = "Tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

}
