package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type linkedinConnection struct {
	firstName      string
	lastName       string
	email          string
	company        string
	position       string
	connectionDate string
}

type invitedPerson struct {
	id          string
	fullName    string
	profileLink string
	firstName   string
	lastName    string
	avatar      string
	title       string
	company     string
	position    string
}

var connections []linkedinConnection
var invitedPeople []invitedPerson

func main() {
	readLinkedinConnections()
	readInvitedFile()

	counter := 0

	for _, invited := range invitedPeople {
		for _, connection := range connections {
			if invited.fullName == connection.getFullName() {
				counter++
			}
		}
	}

	fmt.Println(counter, " invited persons become connections")
}

func (c linkedinConnection) getFullName() string {
	return c.firstName + " " + c.lastName
}

func (c linkedinConnection) printFullName() {
	fmt.Println(c.getFullName())
}

func readLinkedinConnections() {
	linkedinFile, err := os.Open("Connections.csv")

	if err != nil {
		fmt.Println("Error happened while accessing file: ", err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(linkedinFile))
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error happened during reading contents of the file: ", err)
			os.Exit(1)
		}

		connections = append(connections, linkedinConnection{
			firstName:      line[0],
			lastName:       line[1],
			email:          line[2],
			company:        line[3],
			position:       line[4],
			connectionDate: line[5],
		})
	}
}

func readInvitedFile() {
	invitedFile, err := os.Open("Invited.csv")

	if err != nil {
		fmt.Println("Error happened while accessing file: ", err)
		os.Exit(1)
	}

	reader := csv.NewReader(bufio.NewReader(invitedFile))
	reader.LazyQuotes = true

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error happened during reading contents of the file: ", err)
			os.Exit(1)
		}

		invitedPeople = append(invitedPeople, invitedPerson{
			id:          line[0],
			fullName:    line[1],
			profileLink: line[2],
			firstName:   line[3],
			lastName:    line[4],
			avatar:      line[5],
			title:       line[6],
			company:     line[7],
			position:    line[8],
		})
	}
}
