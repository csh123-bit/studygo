package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "고 컨퍼런스"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, int(conferenceTickets), remainingTickets)

	for {
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

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("The slice type: %T\n", bookings)
			fmt.Printf("slice 길이: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. you will 받을것이다 이메일 at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("첫번째 이름은 %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next yaer.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is to short")
			}
			if !isvalidEmail {
				fmt.Println("email address you entered doesn't contain @sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket you entered is invalid")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("welocome to %v booking app\n", confName)
	fmt.Printf("전체 티켓: %v 남은 티켓 %v\n", confTickets, remainingTickets)
	fmt.Println("당신의 표를 얻으세요")
	fmt.Printf("컨퍼런스 티켓은 %T 남은 티켓은 %T, 컨퍼런스 이름은 %T\n", confTickets, remainingTickets, confName)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isvalidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isvalidEmail, isValidTicketNumber
}

//2:13:00
