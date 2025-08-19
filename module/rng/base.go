package rng

import (
	"math/rand"
)

func Rng() int {
	return rand.Intn(6) + 1
}

func Rng2() int {
	return rand.Intn(2) + 1
}

func RngHp() int {
    return rand.Intn(21) + 80
}
