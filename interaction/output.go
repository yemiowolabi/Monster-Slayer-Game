package interaction

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var userName string

type RoundData struct {
	Action              string
	PlayerAttackDamage  int
	PlayerHealValue     int
	MonsterAttackDamage int
	PlayerHealth        int
	MonsterHealth       int
}

func PrintGreeting() {
	fmt.Println("WELCOME TO THE MONSTER SLAYER GAME!!!")
	fmt.Println("Starting a New Game . . .")
	fmt.Print("Please Enter your Name: ")
	userName = GetPlayerName()
	fmt.Printf("%v vs. Monster; LET'S GOOOO!!!\n", userName)
	fmt.Println(strings.Repeat("=", 35))
}

func DispUserActions(specialAttackIsAvailable bool) {
	fmt.Println("Choose Action: ")
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("(1) ATTACK MONSTER")
	fmt.Println("(2) HEAL YOURSELF")
	if specialAttackIsAvailable {
		fmt.Println("(3) SPECIAL ATTACK!")
	}
}

func (roundData *RoundData) ShowRoundStatistics() {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked Monster; DAMAGE: %v\n", roundData.PlayerAttackDamage)
	} else if roundData.Action == "SPECIAL ATTACK" {
		fmt.Printf("Player attacked Monster with SPECIAL ATTACK; DAMAGE: %v\n", roundData.PlayerAttackDamage)
	} else {
		fmt.Printf("Player Healed with: %v\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked Player; DAMAGE: %v\n", roundData.MonsterAttackDamage)

	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)

}

func DeclareWinner(winner string) {
	fmt.Println(strings.Repeat("=", 35))
	fmt.Println("GAME OVER !!!!!!")
	fmt.Println(strings.Repeat("-", 35))

	if winner == "Player" {
		fmt.Printf("YOU(%v) WON !!!\n", userName)
	} else {
		fmt.Println("YOU LOST; The Monster won!")
	}

}

func CreateLogFile(roundSlice *[]RoundData) {
	file, err := os.Create("logfile.txt")
	if err != nil {
		fmt.Println("Creating Log File Failed; EXITING ....")
		log.Fatal(err)
	}
	fmt.Fprintf(file, "PLAYER NAME: ")
	fmt.Fprintf(file, userName+"\n")

	for index, value := range *roundSlice {
		_, err := fmt.Fprintf(file, "Action: "+value.Action+" ; ")
		fmt.Fprintf(file, "Player Attack Damage: "+fmt.Sprint(value.PlayerAttackDamage)+" ; ")
		fmt.Fprintf(file, "Monster Attack Damage: "+fmt.Sprint(value.MonsterAttackDamage)+" ; ")
		fmt.Fprintf(file, "Player Health: "+fmt.Sprint(value.PlayerHealth)+" ; ")
		fmt.Fprintf(file, "Monster Health: "+fmt.Sprint(value.MonsterHealth)+" ; ")
		fmt.Fprintf(file, "Player Heal Value: "+fmt.Sprint(value.PlayerHealValue)+" ; ")
		fmt.Fprintf(file, "Round No. :"+fmt.Sprint(index+1)+"\n")

		if err != nil {
			fmt.Println("Writing Log File Failed; EXITING ....")
			log.Fatal(err)
		}
		defer file.Close()

	}

	fmt.Fprintln(file, strings.Repeat("=", 110))
	fmt.Fprintln(file, "												END OF GAME	                ")
	fmt.Println("Wrote Data to \"logfile.txt\" ")

}
