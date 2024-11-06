package main

import "fmt"

func main() {
	var confrenceName = "PBKK Confrence"
	const confrenceTickets = 100
	var remainingTickets = confrenceTickets

	fmt.Printf("Welcome to %v \n", confrenceName)
	fmt.Printf("Please Book %v as there is only %v out of %v\n", confrenceName, remainingTickets, confrenceTickets)

	var bookings []string

	var userName string
	var bookedTickets int

	for {
		fmt.Printf("Bro, Who are you? [enter name]\n")
		fmt.Scan(&userName)

		fmt.Printf("how many? [enter ticket amount]\n")
		fmt.Scan(&bookedTickets)
		remainingTickets -= bookedTickets
		bookings = append(bookings, userName)

		fmt.Printf("%v has booked %v tickets\n", userName, bookedTickets)
		fmt.Printf("all booker %v", bookings)
	}

}
