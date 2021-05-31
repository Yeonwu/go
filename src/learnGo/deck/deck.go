package deck

import (
	"learnGo/card"
	"math/rand"
	"time"
)

type Deck struct {
	card [52]*card.Card
	top  int
}

func (deck *Deck) push(card card.Card) {
	deck.card[deck.top] = &card
	deck.top += 1
}

func (deck *Deck) pop() *card.Card {
	deck.top -= 1
	return deck.card[deck.top]
}

func (deck *Deck) init() {
	deck.top = 0
	for _, shape := range card.CardShape {
		for _, number := range card.CardNumber {
			card := card.New(shape, number)
			deck.card[deck.top] = card
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

func (deck *Deck) getTop() *card.Card {
	return deck.pop()
}
