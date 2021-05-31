package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Turn struct {
	playerList      []*Player
	dealer          *Dealer
	currentGamerIdx int
}

func (turn *Turn) start(deck *Deck) {
	turn.currentGamerIdx = -1

	deck.suffle()

	var playerNumber int
	fmt.Print("How many players : ")
	fmt.Scan(&playerNumber)
	for i := 0; i < playerNumber; i++ {
		player := new(Player)
		player.chip = 100
		player.draw()

		turn.playerList = append(turn.playerList, player)
	}

	turn.dealer = new(Dealer)
	turn.dealer.draw()
}

func (turn *Turn) play() {
	currentPlayer := turn.playerList[turn.currentGamerIdx]

LOOP:
	for {
		var playerAction rune
		const HIT = 'H'
		const STAY = 'S'
		fmt.Print("Hit or Stay[H/S]")
		fmt.Scanf("\n%c", &playerAction)
		switch playerAction {
		case HIT:
			currentPlayer.hit()
		case STAY:
			currentPlayer.stay()
			break LOOP
		}
	}
}

func (turn *Turn) next() bool {
	turn.currentGamerIdx += 1
	if turn.currentGamerIdx >= len(turn.playerList) {
		return false
	}
	return true
}

func (turn *Turn) end() {
	fmt.Println("Game finished")
}

var cardShape = []string{"Spade", "Diamond", "Heart", "Clover"}
var cardNumber = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

type Card struct {
	shape  string
	number string
}

func (card Card) point() (int, error) {
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

type Deck struct {
	card [52]*Card
	top  int
}

func (deck *Deck) push(card Card) {
	deck.card[deck.top] = &card
	deck.top += 1
}

func (deck *Deck) pop() *Card {
	deck.top -= 1
	return deck.card[deck.top]
}

func (deck *Deck) init() {
	deck.top = 0
	for _, shape := range cardShape {
		for _, number := range cardNumber {
			card := Card{shape, number}
			deck.card[deck.top] = &card
			deck.top += 1
		}
	}
}

func (deck *Deck) suffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.card), func(i, j int) {
		deck.card[i], deck.card[j] = deck.card[j], deck.card[i]
	})
}

func (deck *Deck) getTop() *Card {
	return deck.pop()
}

type Gamer struct {
	cardList []Card
}

func (gamer *Gamer) draw() {
	fmt.Println("Drawing Cards...")
}

type Dealer struct {
	Gamer
}

func (dealer *Dealer) play() {

}

type Player struct {
	Gamer
	chip int
}

func (player *Player) hit() {
	fmt.Println("Chose to hit")
}

func (player *Player) stay() {
	fmt.Println("Chose to stay")
}

func main() {
	turn := new(Turn)
	deck := new(Deck)

	deck.init()

	turn.start(deck)

	for turn.next() {
		turn.play()
	}

	turn.dealer.play()
	turn.end()
}
