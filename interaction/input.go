package interaction

import (
	"bufio"
	"fmt"

	//	"log"
	"os"
	//	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerName() string {
	inputRead, _ := reader.ReadString('\n')
	inputRead = strings.Trim(inputRead, "\n\r")
	return inputRead
}
func GetPlayerChoice(specialAttackIsAvailable bool) string {
	for {
		playerChoice, _ := getPlayerInput()
		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAttackIsAvailable {
			return "SPECIAL ATTACK"
		} else {
			fmt.Println("Error Fetching User Input; Please try again!")
		}
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your Choice: ")
	inputRead, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	inputRead = strings.Trim(inputRead, "\n\r")
	return inputRead, nil
}
