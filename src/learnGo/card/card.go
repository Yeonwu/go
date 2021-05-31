package card

import (
	"errors"
	"strconv"
)

var CardShape = []string{"Spade", "Diamond", "Heart", "Clover"}
var CardNumber = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

type Card struct {
	shape  string
	number string
}

func (card Card) Point() (int, error) {
	switch card.number {
	case "2", "3", "4", "5", "6", "7", "8", "9", "10":
		r, _ := strconv.Atoi(card.number)
		return r, nil
	case "J", "Q", "K":
		return 10, nil
	case "A":
		return 1, nil
	}
	return -1, errors.New("Something wrong with Card.point")
}

func New(inShape string, inNumber string) *Card {
	return &Card{shape: inShape, number: inNumber}
}
