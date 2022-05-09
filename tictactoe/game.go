package tictactoe

import (
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Turn int

const (
	Circle Turn = 0
	Cross  Turn = 1
)

type Game struct {
	images *Images

	turn Turn
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
		turn: Circle,
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
		mx, my := ebiten.CursorPosition()
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(mx)-67, float64(my)-67)
		switch g.turn {
		case Circle:
			g.images.game.DrawImage(g.images.circle, opts)
		case Cross:
			g.images.game.DrawImage(g.images.cross, opts)
		}
		g.turn ^= 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.images.board, nil)
	screen.DrawImage(g.images.game, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WindowWidth, WindowHight
}
