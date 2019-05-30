package main

import (
	"image/color"
	_ "image/png"
	"log"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)


type Game struct {
	//Player's position
	positionx float64
	positiony float64
	vy float64 

	// Camera
	cameraX int
	cameraY int
}

func NewGame() *Game {
	g := &Game{}
	g.in()
	return g
}

func (g *Game) in() {
	g.positionx = 40
	g.positiony = 2600
	g.cameraX = 0
	g.cameraY = 0
}

func (g *Game) drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	h := 16
	op.GeoM.Translate(g.positionx, -float64(h)/2.0)
	op.GeoM.Translate(g.positionx, float64(h)/2.0)
	op.GeoM.Translate(g.positionx, float64(g.positiony/16.0))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(player, op)
}

//setting game elements
var player *ebiten.Image
var floor *ebiten.Image
var i float64


//variable used to make the jump mutual exclusive
var jump_count = 0

//opts is a global variable because it would not update correctly inside the "update" func
var opts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}

// update is called every frame (1/60 [s]).
func (g *Game) update(screen *ebiten.Image) error {
    
	screen.Fill(color.Black)
    
	//only for debugging
	fmt.Println(g.positionx, g.positiony)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %0.2f", ebiten.CurrentFPS()))
    
	var err error
	//creating floor
	floor, _, err = ebitenutil.NewImageFromFile("floor.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
    
	//creating player
	if player == nil {
		player, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
	}
    
	player.Fill(color.White)
    
	//controls and screen borders
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.positionx < 101{
		g.positionx += 1
		opts.GeoM.Translate(1, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.positionx > 0{
		g.positionx -= 1
		opts.GeoM.Translate(-1, 0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && jump_count == 0 && g.positiony == 2630{
		jump_count = 1
		g.vy = -96
		jump_count = 0
	}
	// Gravity
	g.vy += 4
	if g.vy > 96{
		g.vy = 96
	}
	g.positiony += g.vy
	if g.positiony > 2630{
		g.positiony = 2630
	}
    
	//defining floor image options
	optf := &ebiten.DrawImageOptions{}
    
	optf.GeoM.Translate(0,180)
    
	screen.DrawImage(floor, optf)
	g.drawPlayer(screen)
	return nil
}

func main() {
	g := NewGame()
	if err := ebiten.Run(g.update, 320, 240, 2, "TankRun"); err != nil {
		panic(err)
	}
}
