package main

import "fmt"

func main(){

	var conferenceName = "Blaise Conferences"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Print("================================\n")
	fmt.Println("Welcome to ", conferenceName ," the booking app")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend.")
	fmt.Print("================================\n")
	
}
