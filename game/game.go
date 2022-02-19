package game

import "fmt"

func PlayGame(hand Hand) {
	opponentHand := FindRandomHand()
	fmt.Println("Oponent Hand: " + opponentHand)
	if IsWinner(hand, opponentHand) {
		fmt.Println("You win")
	} else if IsWinner(opponentHand, hand) {
		fmt.Println("You loose")
	} else {
		fmt.Println("Drawn round")
	}
}
