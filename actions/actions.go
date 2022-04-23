package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

var presentPlayerHealth int = 200
var presentMonsterHealth int = 200

func AttackMonster(specialAttackIsAvailable bool) int {
	minAttackVal := 10
	maxAttackVal := 20
	if specialAttackIsAvailable {
		minAttackVal = 20
		maxAttackVal = 40
	}

	damageValue := genRandBetween(minAttackVal, maxAttackVal)
	presentMonsterHealth -= damageValue
	return damageValue
}

func HealPlayer() int {
	minHealVal := 15
	maxHealVal := 25

	healValue := genRandBetween(minHealVal, maxHealVal)

	healthDifference := 200 - presentPlayerHealth
	if healthDifference >= healValue {
		presentPlayerHealth += healValue
		return healValue

	} else {
		presentPlayerHealth = 200
		return healthDifference
	}
}

func AttackPlayer() int {
	minAttackVal := 15
	maxAttackVal := 25

	damageValue := genRandBetween(minAttackVal, maxAttackVal)
	presentPlayerHealth -= damageValue
	return damageValue
}

func UpdateHealth() (int, int) {
	return presentPlayerHealth, presentMonsterHealth
}

func genRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min //Generates random no. between min and max
}
