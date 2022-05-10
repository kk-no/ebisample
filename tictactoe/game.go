package tictactoe

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TurnPlayer int

const (
	Circle TurnPlayer = 0
	Cross  TurnPlayer = 1
)

func (t TurnPlayer) String() string {
	switch t {
	case Circle:
		return "○"
	case Cross:
		return "×"
	default:
		return ""
	}
}

type Game struct {
	images     *Images
	board      [3][3]string
	turnPlayer TurnPlayer
}

type Images struct {
	game   *ebiten.Image
	board  *ebiten.Image
	circle *ebiten.Image
	cross  *ebiten.Image
}

func New() (*Game, error) {
	board, err := loadImage("images/board.png")
	if err != nil {
		return nil, err
	}
	circle, err := loadImage("images/circle.png")
	if err != nil {
		return nil, err
	}
	cross, err := loadImage("images/cross.png")
	if err != nil {
		return nil, err
	}
	return &Game{
		images: &Images{
			game:   ebiten.NewImage(WindowWidth, WindowHight),
			board:  ebiten.NewImageFromImage(board),
			circle: ebiten.NewImageFromImage(circle),
			cross:  ebiten.NewImageFromImage(cross),
		},
		turnPlayer: Circle,
	}, nil
}

func loadImage(name string) (image.Image, error) {
	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return png.Decode(bytes.NewReader(b))
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.DrawSymbol(x, y)
		g.SetBoardSymbol(x, y)

		if winner := g.CheckWinner(); winner != "" {
			fmt.Printf("%s win!\n", winner)
			return nil
		}
		g.turnPlayer ^= 1 // XOR
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.images.board, nil)
	screen.DrawImage(g.images.game, nil)
}

func (g *Game) DrawSymbol(x, y int) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(x-67), float64(y-67))
	switch g.turnPlayer {
	case Circle:
		g.images.game.DrawImage(g.images.circle, opts)
	case Cross:
		g.images.game.DrawImage(g.images.cross, opts)
	}
}

func (g *Game) SetBoardSymbol(x, y int) {
	g.board[x/160][y/160] = g.turnPlayer.String()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHight
}

func (g *Game) CheckWinner() string {
	for i := range g.board {
		if IsEqual(g.board[i][0], g.board[i][1], g.board[i][2]) {
			return g.turnPlayer.String()
		}
		if IsEqual(g.board[0][i], g.board[1][i], g.board[2][i]) {
			return g.turnPlayer.String()
		}
	}
	if IsEqual(g.board[0][0], g.board[1][1], g.board[2][2]) {
		return g.turnPlayer.String()
	}
	if IsEqual(g.board[0][2], g.board[1][1], g.board[2][0]) {
		return g.turnPlayer.String()
	}
	return ""
}

func IsEqual(s1, s2, s3 string) bool {
	return s1 != "" && s1 == s2 && s2 == s3
}
