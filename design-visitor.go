package main

import "fmt"

type PlayerType int

const (
	Beginner PlayerType = iota
	Pro
)

// visitor interface
type ScorerInterface interface {
	AddScore(player *Player) int
}

// visitor implementation
type BasicScorer struct {
}

// visitor logic, caller is passed by reference
func (sc BasicScorer) AddScore(player *Player) int {
	if player.playerType == Beginner {
		player.score = 10
	} else if player.playerType == Pro {
		player.score = 50
	}
	return 0
}

// caller interface
type PlayerInteface interface {
	AddScore(sc ScorerInterface)
}

// caller implentation
type Player struct {
	playerType PlayerType
	name       string
	score      int
}

// caller method with vistor as an argument
func (p *Player) AddScore(sc ScorerInterface) {
	sc.AddScore(p)
}

func main() {
	player1 := Player{name: "Jack", playerType: Beginner}
	player2 := Player{name: "Lisa", playerType: Pro}

	scorer := BasicScorer{}

	// scorer logic is decoupled from the caller
	player1.AddScore(scorer)
	player2.AddScore(scorer)

	fmt.Printf("player 1 score = %d\n", player1.score)
	fmt.Printf("player 2 score = %d\n", player2.score)
}
