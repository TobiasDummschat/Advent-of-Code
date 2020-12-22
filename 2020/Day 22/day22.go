package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Deck []int

func (deck *Deck) draw() int {
	card := (*deck)[0]
	*deck = (*deck)[1:]
	return card
}

func (deck *Deck) placeAtBottom(card int) {
	*deck = append(*deck, card)
}

func equals(deck1, deck2 Deck) bool {
	if len(deck1) != len(deck2) {
		return false
	}
	for i := 0; i < len(deck1); i++ {
		if deck1[i] != deck2[i] {
			return false
		}
	}
	return true
}

type Memory [][2]Deck

func (memory Memory) happenedBefore(deck1, deck2 Deck) bool {
	for _, setup := range memory {
		pastDeck1, pastDeck2 := setup[0], setup[1]
		if equals(deck1, pastDeck1) && equals(deck2, pastDeck2) {
			return true
		}
	}
	return false
}

type Counter struct {
	round, game, highestGame int
}

func newCounter() Counter {
	return Counter{1, 1, 1}
}

func main() {
	deck1, deck2 := parseInput("2020\\Day 22\\day22_input")

	counter := newCounter()
	winner, winningScore := combatGame(deck1, deck2, false, false, counter)
	fmt.Printf("Player %d wins the game of Combat with a score of %d.\n", winner, winningScore)

	counter = newCounter()
	winner, winningScore = combatGame(deck1, deck2, true, false, counter)
	fmt.Printf("Player %d wins the game of Recursive Combat with a score of %d.\n", winner, winningScore)
}

func combatGame(deck1, deck2 Deck, recursive, log bool, counter Counter) (winner, winningScore int) {
	memory := Memory{}

	if recursive && log {
		fmt.Printf("=== Game %d ===\n\n")
	}

	// play the game until it is over
	for len(deck1) > 0 && len(deck2) > 0 {
		// check if setup happened before this game, otherwise remember setup
		if recursive {
			if memory.happenedBefore(deck1, deck2) {
				return 1, score(deck1)
			}
			memory = append(memory, [2]Deck{deck1, deck2})
		}

		// play a round
		combatRound(&deck1, &deck2, recursive, log, &counter)
		counter.round++
	}

	// game is over
	if len(deck1) == 0 { // player 2 wins
		if log {
			fmt.Printf("The winner of game %d is player 2!\n\n", counter.game)
		}
		return 2, score(deck2)
	} else { // player 1 wins
		if log {
			fmt.Printf("The winner of game %d is player 1!\n\n", counter.game)
		}
		return 1, score(deck1)
	}
}

func combatRound(deck1, deck2 *Deck, recursive, log bool, counter *Counter) {
	if log {
		if recursive {
			fmt.Printf("-- Round %d (Game %d) --\n", counter.round, counter.game)
		} else {
			fmt.Printf("-- Round %d --\n", counter.round)
		}
		fmt.Printf("Player 1's deck: %v\n", *deck1)
		fmt.Printf("Player 2's deck: %v\n", *deck2)
	}

	card1 := deck1.draw()
	card2 := deck2.draw()

	if log {
		fmt.Printf("Player 1 plays: %d\n", card1)
		fmt.Printf("Player 2 plays: %d\n", card2)
	}

	winner := -1

	if recursive && card1 <= len(*deck1) && card2 <= len(*deck2) {
		// play sub-game
		// copy top cards
		topCards1, topCards2 := Deck{}, Deck{}
		for _, card := range (*deck1)[:card1] {
			topCards1.placeAtBottom(card)
		}
		for _, card := range (*deck2)[:card2] {
			topCards2.placeAtBottom(card)
		}

		if log {
			fmt.Printf("Playing sub-game to determine the winner...\n\n")
		}
		subGameCounter := Counter{
			round:       1,
			game:        counter.highestGame + 1,
			highestGame: counter.highestGame + 1,
		}
		counter.highestGame = subGameCounter.game
		winner, _ = combatGame(topCards1, topCards2, true, log, subGameCounter)
		if log {
			fmt.Printf("...anyway, back to game %d.\n", counter.game)
		}
	} else if card1 > card2 { // player 1 wins round
		winner = 1
	} else if card2 > card1 { // player 2 wins round
		winner = 2
	} else {
		panic(fmt.Errorf("round is a draw"))
	}
	if winner == 1 {
		deck1.placeAtBottom(card1)
		deck1.placeAtBottom(card2)
	} else if winner == 2 {
		deck2.placeAtBottom(card2)
		deck2.placeAtBottom(card1)
	} else {
		panic(fmt.Errorf("game is over without a winner"))
	}

	if log {
		if recursive {
			fmt.Printf("Player %d wins round %d of game %d!\n\n", winner, counter.round, counter.game)
		} else {
			fmt.Printf("Player %d wins the round!\n\n", winner)
		}
	}
}

func score(deck Deck) int {
	score := 0
	size := len(deck)
	for i, card := range deck {
		score += card * (size - i)
	}
	return score
}

func parseInput(path string) (deck1, deck2 Deck) {
	contents, _ := ioutil.ReadFile(path)
	players := strings.Split(string(contents), "\r\n\r\n")
	player1, player2 := players[0], players[1]
	deck1 = parseDeck(player1)
	deck2 = parseDeck(player2)
	return deck1, deck2
}

func parseDeck(player string) Deck {
	deck := Deck{}
	lines := strings.Split(player, "\r\n")
	cardStrings := lines[1:]
	for _, cardString := range cardStrings {
		card, _ := strconv.Atoi(cardString)
		deck.placeAtBottom(card)
	}
	return deck
}
