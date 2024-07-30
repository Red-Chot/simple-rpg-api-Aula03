package entity

import (
	"math/rand"
	"time"

	  "github.com/google/uuid"
)





type Enemy struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Life     int    `json:"life"`
	Attack   int    `json:"attack"`
}



func NewEnemy(nickname string) *Enemy {
	rand.Seed(time.Now().UnixNano())
	return &Enemy{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     rand.Intn(10) + 1,
		Attack:   rand.Intn(10) + 1,
	}
}

















