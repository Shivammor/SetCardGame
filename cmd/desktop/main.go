package main

import (
    "log"
    "setcardgame/internal/graphics"
    "setcardgame/internal/scene"

    "github.com/hajimehoshi/ebiten/v2"
)

func main() {
    // Load assets from files
    bg := graphics.LoadImageFromFile("./assets/background.png")
    btnNormalImg := graphics.LoadImageFromFile("./assets/UI_elemets/Default@4x.png")
    btnHoverImg := graphics.LoadImageFromFile("./assets/UI_elemets/Hover@4x.png")
    squareNormalImg := graphics.LoadImageFromFile("./assets/UI_elemets/Default@4x_square.png")
    squareHoverImg := graphics.LoadImageFromFile("./assets/UI_elemets/Hover@4x_square.png")
    fontSrc := graphics.LoadFontFromFile("./assets/FSEX300.ttf")

    // Create menu scene
    menuScene := scene.NewMenuScene(bg, btnNormalImg, btnHoverImg, squareNormalImg, squareHoverImg, fontSrc)

    ebiten.SetWindowSize(scene.ScreenWidth, scene.ScreenHeight)
    ebiten.SetWindowTitle("Set Card Game")

    if err := ebiten.RunGame(menuScene); err != nil {
        log.Fatal(err)
    }
}

