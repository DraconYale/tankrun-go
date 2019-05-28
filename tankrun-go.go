package main

import (
    "image/color"
    _ "image/png"
    "log"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/hajimehoshi/ebiten/inpututil"
    
)

var player *ebiten.Image
var floor *ebiten.Image

var(
    positionx float64
    positiony float64
)

var opts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}


// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {
    
    screen.Fill(color.Black)
    
    var err error
	floor, _, err = ebitenutil.NewImageFromFile("floor.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
    
    if player == nil {
        player, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
    }
    
    player.Fill(color.White)
    
    if ebiten.IsKeyPressed(ebiten.KeyRight){
        opts.GeoM.Translate(positionx+(float64(inpututil.KeyPressDuration(ebiten.KeyRight))), 0)
    }
    if ebiten.IsKeyPressed(ebiten.KeyLeft){
        opts.GeoM.Translate(positionx-(float64(inpututil.KeyPressDuration(ebiten.KeyLeft))), 0)
    }
    if ebiten.IsKeyPressed(ebiten.KeyUp){
        opts.GeoM.Translate(0, positiony-(float64(inpututil.KeyPressDuration(ebiten.KeyUp))))
    }
    if ebiten.IsKeyPressed(ebiten.KeyDown){
        opts.GeoM.Translate(0, positiony+(float64(inpututil.KeyPressDuration(ebiten.KeyDown))))
    }
    optf := &ebiten.DrawImageOptions{}
    optf.GeoM.Translate(0,180)
    screen.DrawImage(floor, optf)
    screen.DrawImage(player, opts)
    
    return nil
}

func main() {
    // Call ebiten.Run to start your game loop.
    if err := ebiten.Run(update, 320, 240, 2, "TankRun"); err != nil {
        panic(err)
    }
}
