package main

import (
	"bufio"
	"fmt"
	"hello-world-go/game"
	"os"
	"strings"
)

func init() {
}

func main() {
	fmt.Println("Welcome to Rock/Paper/Scissor game.")
	fmt.Println("Would you like to play? (Y/N).")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if strings.TrimSuffix(text, "\r\n") == "Y" {
		fmt.Println("Playing Game")
		fmt.Println("Choose your Hand: (Rock/Paper/Scissor)")
		hand, _ := reader.ReadString('\n')
		handEnum, error := game.FindHand(hand)

		if error != nil {
			fmt.Println("Hand not found")
			return
		}

		game.PlayGame(handEnum)
	}
}
