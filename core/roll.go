package core

import "math/rand"

var (
	Version   string
	BuildDate string
)

// RollDie rolls a single `sides` side die.
func RollDie(sides int) int {
	return rand.Intn(sides) + 1
}
