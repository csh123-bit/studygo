package main

import (
	"fmt"
	"golang/helper"
	"sync"
	"time"
)

var conferenceName = "고 컨퍼런스"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("첫번째 이름은 %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next yaer.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is to short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of ticket you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("welocome to %v booking app\n", conferenceName)
	fmt.Printf("전체 티켓: %v 남은 티켓 %v\n", conferenceTickets, remainingTickets)
	fmt.Println("당신의 표를 얻으세요")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("당신의 이름을 입력하세요")
	fmt.Scan(&firstName)
	fmt.Println("당신의 성을을 입력하세요")
	fmt.Scan(&lastName)
	fmt.Println("당신의 이메일을 입력하세요")
	fmt.Scan(&email)
	fmt.Println("티켓을 입력하세요")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will 받을것이다 이메일 at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("--------------")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("--------------")
	wg.Done()
}

//3:06:39
