// 1. Despot some money
// 2. Determine number of lines to bet on
// 3. Collect a bet amount
// 4. Spin the slot machine
// 5. check if the user won
// 6. give the user their winnings
// 7. play again

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const ROWS = 3
const COLS = 3

var depositAmount string
var lines string
var bet string

var symbolCount map[string]int
var symbolValues map[string]int

// 1. Despot some money
func deposit() float64 {
	for {
		fmt.Printf("Enter a deposit amount: $")
		fmt.Scanln(&depositAmount)

		numberDepositAmount, err := strconv.ParseFloat(depositAmount, 64)
		if err != nil || numberDepositAmount <= 0 {
			fmt.Println("Invalid deposit amount, try again.")
		} else {
			return numberDepositAmount
		}
	}
}

// 2. Determine number of lines to bet on
func getNumberOfLines() float64 {
	for {
		fmt.Printf("Enter the number of lines to ben on (1-3): ")
		fmt.Scanln(&lines)

		numberOfLines, err := strconv.ParseFloat(lines, 64)
		if err != nil || numberOfLines <= 0 || numberOfLines > 3 {
			fmt.Println("Invalid number of lines, try again.")
		} else {
			return numberOfLines
		}
	}
}

// 3. Collect a bet amount
func getBet(balance, lines float64) float64 {
	fmt.Printf("Enter the bet per line: ")
	fmt.Scanln(&bet)

	numberBet, err := strconv.ParseFloat(bet, 64)
	if err != nil || numberBet <= 0 || numberBet > balance/lines {
		fmt.Println("Invalid bet, try again.")
	}

	return numberBet
}

// 4. Spin the slot machine
func machineSpin() [][]string {
	var collectAllSymbols []string
	for symbol, count := range symbolCount {
		for i := 0; i < count; i++ {
			collectAllSymbols = append(collectAllSymbols, symbol)
		}
	}

	reels := make([][]string, ROWS)
	reelSymbols := make([]string, len(collectAllSymbols))

	for k := 0; k < COLS; k++ {
		reels[k] = make([]string, ROWS)
		copy(reelSymbols, collectAllSymbols)
		for j := 0; j < ROWS; j++ {
			randomIndex := rand.Intn(len(reelSymbols))
			selectSymbol := reelSymbols[randomIndex]
			reels[k][j] = selectSymbol
			reelSymbols = append(reelSymbols[:randomIndex], reelSymbols[randomIndex+1:]...)
		}
	}
	return reels
}

// Group all first,second and third elements to their specific slices
// Example - [[A, B ,C],[C,A,D],[D,B,A]] to [[A,C,D] [B,A,C] [C,D,A]]
func transpose(reels [][]string) [][]string {
	rows := make([][]string, ROWS)

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			rows[i] = append(rows[i], reels[j][i])
		}
	}
	return rows
}

// Print out the Rows
func printRows(rows [][]string) {
	for _, value := range rows {

		rowString := ""
		for i, symbol := range value {
			rowString += symbol
			if i != (len(value) - 1) {
				rowString += " | "
			}
		}
		fmt.Println(rowString)
	}
}

// 5. check if the user won
func getWinnings(rows [][]string, bet float64, lines float64) float64 {
	winnings := 0.0

	for row := 0; row < int(lines); row++ {
		symbols := rows[row]
		allSame := true

		for _, symbol := range symbols {
			if symbol != symbols[0] {
				allSame = false
				break
			}
		}

		if allSame {
			winnings += bet * float64(symbolValues[symbols[0]])
		}
	}

	return winnings
}

func game() {
	symbolCount = make(map[string]int)
	symbolCount["A"] = 2
	symbolCount["B"] = 4
	symbolCount["C"] = 6
	symbolCount["D"] = 8

	symbolValues = make(map[string]int)
	symbolValues["A"] = 10
	symbolValues["B"] = 9
	symbolValues["C"] = 8
	symbolValues["D"] = 7

	balance := deposit()
	for {
		fmt.Printf("You have balance of $ %2.f\n", balance)
		linesNumber := getNumberOfLines()
		bet := getBet(balance, linesNumber)
		balance -= bet * linesNumber
		reels := machineSpin()
		rows := transpose(reels)
		printRows(rows)
		winnings := getWinnings(rows, bet, linesNumber)
		balance += winnings
		fmt.Printf("You won, $ %2.f\n", winnings)

		if balance <= 0 {
			fmt.Println("You ran out of money!")
			break
		}

		playAgain := ""
		fmt.Printf("Do you want to play again (y/n)? ")
		_, err := fmt.Scanln(&playAgain)

		if err != nil {
			fmt.Println("You have enter invalid choice. Please selected either y or n.")
		}

		if playAgain != "y" {
			break
		}
	}

}

func main() {
	game()
}
