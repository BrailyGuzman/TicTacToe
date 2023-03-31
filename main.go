package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type TicTacToe struct {
	table    [9]string
	player   string
	computer string
}

/*
Tic Tac Toe Table

| X || * || * |
| * || * || * |
| * || * || * |
*/

func main() {
	ClearConsole()
	game := TicTacToe{}
	game.player = color.GreenString("X")
	game.computer = color.RedString("O")
	game.table = [9]string{"*", "*", "*", "*", "*", "*", "*", "*", "*"}

	for {
		PrintTable(&game)
		fmt.Print("Enter your move (1-9): ")
		var userMove int
		fmt.Scanln(&userMove)

		playerMove(&game, userMove)
		if CheckDraw(&game) {
			fmt.Println("You Drawed!")
			break
		}

		if CheckWin(&game, game.player) {
			fmt.Println("You won!")
			break
		}

		ComputerMove(&game)

		if CheckDraw(&game) {
			fmt.Println("Computer Drawed!")
			break
		}
		if CheckWin(&game, game.computer) {
			fmt.Println("You lost! :(")
			break
		}

		ClearConsole()
	}

	PrintTable(&game)
	Reset(&game)
}

func PrintTable(t *TicTacToe) {

	for i := 0; i < 3; i++ {
		fmt.Print("| ", t.table[i], " |")
	}
	fmt.Println()
	for i := 3; i < 6; i++ {
		fmt.Print("| ", t.table[i], " |")
	}
	fmt.Println()
	for i := 6; i < 9; i++ {
		fmt.Print("| ", t.table[i], " |")
	}
	fmt.Println()
}

func playerMove(t *TicTacToe, userMove int) {
	userMove--
	if t.table[userMove] != "*" {
		fmt.Println("That cell is already taken. Please choose another cell.")
		return
	}

	t.table[userMove] = t.player
}

func ComputerMove(t *TicTacToe) {
	rand.Seed(time.Now().UnixNano())

	for {
		computerMove := rand.Intn(9)

		if t.table[computerMove] == "*" {
			t.table[computerMove] = t.computer
			break
		}
	}
}

func CheckWin(t *TicTacToe, player string) bool {

	// check rows
	if t.table[0] == player && t.table[1] == player && t.table[2] == player {
		return true
	}
	if t.table[3] == player && t.table[4] == player && t.table[5] == player {
		return true
	}
	if t.table[6] == player && t.table[7] == player && t.table[8] == player {
		return true
	}

	// check columns
	if t.table[0] == player && t.table[3] == player && t.table[6] == player {
		return true
	}

	if t.table[1] == player && t.table[4] == player && t.table[7] == player {
		return true
	}

	if t.table[2] == player && t.table[5] == player && t.table[8] == player {
		return true
	}

	// check diagonals
	if t.table[0] == player && t.table[4] == player && t.table[8] == player {
		return true
	}

	if t.table[2] == player && t.table[4] == player && t.table[6] == player {
		return true
	}

	return false
}

func CheckDraw(t *TicTacToe) bool {
	for i := 0; i < 9; i++ {
		if t.table[i] == "*" {
			return false
		}
	}
	return true
}

func Reset(t *TicTacToe) {
	for i := 0; i < 9; i++ {
		t.table[i] = "*"
	}
}

func ClearConsole() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
