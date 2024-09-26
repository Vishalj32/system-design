package service

import "math/rand"

func RollDice(noOfDice int) int {
	return rand.Intn(6*noOfDice) + 1
}
