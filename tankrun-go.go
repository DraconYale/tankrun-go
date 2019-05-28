package main

import (
    "image/color"
    _ "image/png"
    "log"
    "fmt"
    "time"
    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/hajimehoshi/ebiten/inpututil"
    
)

//setting game elements
var player *ebiten.Image
var floor *ebiten.Image
var i float64
//setting player position
var(
    positionx float64 = 40.0
    positiony float64 = 164.0
)

//opts is a global variable because it would not update correctly inside the "update" func
var opts *ebiten.DrawImageOptions = &ebiten.DrawImageOptions{}

func jump(){
    for t:=time.Now(); time.Since(t)<(time.Second); {
        if(positiony>140){
            positiony -= 0.0001
            opts.GeoM.Translate(0, -0.0001)
        }
    }
    for t:=time.Now(); time.Since(t)<(time.Second); {
        if(positiony<164){
            positiony += 0.1
            opts.GeoM.Translate(0, 0.1)
        }
    }
    
}

// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {
    
    screen.Fill(color.Black)
    
    //only for debugging
    fmt.Println(positionx, positiony)
    
    var err error
	//creating floor
	floor, _, err = ebitenutil.NewImageFromFile("floor.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
    
    //creating player
    if player == nil {
        player, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
        opts.GeoM.Translate(40,164)
    }
    
    player.Fill(color.White)
    
    //controls and screen borders
    if ebiten.IsKeyPressed(ebiten.KeyRight) && positionx < 304{
        positionx += 2
        opts.GeoM.Translate(2, 0)
    }
    if ebiten.IsKeyPressed(ebiten.KeyLeft) && positionx > 0{
        positionx -= 2
        opts.GeoM.Translate(-2, 0)
    }
    if inpututil.IsKeyJustPressed(ebiten.KeySpace) && positiony > 0{
        go jump()
    }
    
    //defining floor image options
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
