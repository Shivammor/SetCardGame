package main

import (
    "log"
    "setcardgame/internal/assets"
    "setcardgame/internal/graphics"
    "setcardgame/internal/scene"

    "github.com/hajimehoshi/ebiten/v2"
)

func main() {
    // Load assets from embedded data
    bg := graphics.LoadImageFromBytes(assets.BackgroundPNG)
    btnNormalImg := graphics.LoadImageFromBytes(assets.DefaultButtonPNG)
    btnHoverImg := graphics.LoadImageFromBytes(assets.HoverButtonPNG)
    squareNormalImg := graphics.LoadImageFromBytes(assets.SquareDefaultPNG)
    squareHoverImg := graphics.LoadImageFromBytes(assets.SquareHoverPNG)
    fontSrc := graphics.LoadFontFromBytes(assets.FontTTF)

    // Create menu scene
    menuScene := scene.NewMenuScene(bg, btnNormalImg, btnHoverImg, squareNormalImg, squareHoverImg, fontSrc)

    ebiten.SetWindowSize(scene.ScreenWidth, scene.ScreenHeight)
    ebiten.SetWindowTitle("Set Card Game")

    if err := ebiten.RunGame(menuScene); err != nil {
        log.Fatal(err)
    }
}

