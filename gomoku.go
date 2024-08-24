package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	boardSize = 15
)

type Gomoku struct {
	window  fyne.Window
	buttons [][]*widget.Button
	board   [][]int
	player  int
}

func newGomoku() *Gomoku {
	g := &Gomoku{
		buttons: make([][]*widget.Button, boardSize),
		board:   make([][]int, boardSize),
		player:  1,
	}

	for i := range g.buttons {
		g.buttons[i] = make([]*widget.Button, boardSize)
		g.board[i] = make([]int, boardSize)
	}

	return g
}

func (g *Gomoku) createUI() {
	grid := container.NewGridWithColumns(boardSize)

	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			button := widget.NewButton("", g.makeMove(x, y))
			g.buttons[y][x] = button
			grid.Add(button)
		}
	}

	resetButton := widget.NewButton("Reset", func() {
		g.resetGame()
	})

	content := container.NewVBox(grid, resetButton)
	g.window.SetContent(content)
}

func (g *Gomoku) makeMove(x, y int) func() {
	return func() {
		if g.board[y][x] == 0 {
			g.board[y][x] = g.player
			g.buttons[y][x].SetText(g.playerSymbol())
			g.buttons[y][x].Disable()

			if g.checkWin(x, y) {
				g.endGame(g.player)
			} else {
				g.player = 3 - g.player // Switch player (1 -> 2, 2 -> 1)
			}
		}
	}
}

func (g *Gomoku) playerSymbol() string {
	if g.player == 1 {
		return "X"
	}
	return "O"
}

func (g *Gomoku) checkWin(x, y int) bool {
	// Implement win checking logic here
	// This is a simplified version, you may want to improve it
	directions := [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}
	for _, dir := range directions {
		count := 1
		for i := 1; i < 5; i++ {
			nx, ny := x+dir[0]*i, y+dir[1]*i
			if nx < 0 || nx >= boardSize || ny < 0 || ny >= boardSize || g.board[ny][nx] != g.player {
				break
			}
			count++
		}
		for i := 1; i < 5; i++ {
			nx, ny := x-dir[0]*i, y-dir[1]*i
			if nx < 0 || nx >= boardSize || ny < 0 || ny >= boardSize || g.board[ny][nx] != g.player {
				break
			}
			count++
		}
		if count >= 5 {
			return true
		}
	}
	return false
}

func (g *Gomoku) endGame(winner int) {
	for i := range g.buttons {
		for j := range g.buttons[i] {
			g.buttons[i][j].Disable()
		}
	}
	g.window.SetTitle(fmt.Sprintf("Player %d wins!", winner))
}

func (g *Gomoku) resetGame() {
	g.board = make([][]int, boardSize)
	for i := range g.board {
		g.board[i] = make([]int, boardSize)
	}
	g.player = 1
	for y := 0; y < boardSize; y++ {
		for x := 0; x < boardSize; x++ {
			g.buttons[y][x].SetText("")
			g.buttons[y][x].Enable()
		}
	}
	g.window.SetTitle("Gomoku")
}

func main() {
	a := app.New()
	w := a.NewWindow("Gomoku")
	w.Resize(fyne.NewSize(400, 430))

	game := newGomoku()
	game.window = w
	game.createUI()

	w.ShowAndRun()
}