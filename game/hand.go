package game

import (
	"errors"
	"strings"
	"time"
)

type Hand string

const (
	Rock    = "ROCK"
	Paper   = "PAPER"
	Scissor = "SCISSOR"
)

var values = []Hand{
	Rock, Paper, Scissor,
}

var winners map[Hand]Hand = map[Hand]Hand{Rock: Scissor, Scissor: Paper, Paper: Rock}

func FindHand(value string) (Hand, error) {
	upper := strings.TrimSpace(strings.ToUpper(value))
	var handSelected Hand = ""
	for _, hand := range values {
		if hand == Hand(upper) {
			handSelected = hand
		}
	}
	if handSelected == "" {
		return "", errors.New("Hand not found")
	}
	return handSelected, nil
}

func FindRandomHand() Hand {
	position := time.Now().UnixMilli() % 3
	return values[position]
}

func IsWinner(hand1 Hand, hand2 Hand) bool {
	return winners[hand1] == hand2
}
