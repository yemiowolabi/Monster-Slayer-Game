package main

import (
	"monster_slayer_game/actions"
	"monster_slayer_game/interaction"
)

var presentRound = 0
var roundSlice = []interaction.RoundData{}

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""(No Winner Yet)

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string { //this function should return if there was a winner yet
	presentRound++
	var isSpecialRound bool
	if presentRound%3 == 0 {
		isSpecialRound = true
	} else {
		isSpecialRound = false
	}
	interaction.DispUserActions(isSpecialRound)
	playerChoice := interaction.GetPlayerChoice(isSpecialRound)
	var playerHealth int
	var monsterHealth int
	var playerAttackDamage int
	var playerHealValue int
	var monsterAttackDamage int

	if playerChoice == "ATTACK" {
		playerAttackDamage = actions.AttackMonster(false)
	} else if playerChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDamage = actions.AttackMonster(true)
	}
	monsterAttackDamage = actions.AttackPlayer()
	playerHealth, monsterHealth = actions.UpdateHealth()

	roundData := interaction.RoundData{
		Action:              playerChoice,
		PlayerHealth:        playerHealth,
		MonsterHealth:       monsterHealth,
		PlayerHealValue:     playerHealValue,
		PlayerAttackDamage:  playerAttackDamage,
		MonsterAttackDamage: monsterAttackDamage,
	}
	(&roundData).ShowRoundStatistics()

	roundSlice = append(roundSlice, roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	} else {
		return ""
	}
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
	interaction.CreateLogFile(&roundSlice)
}
