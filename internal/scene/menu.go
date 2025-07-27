package scene

import (
    "log"
    "setcardgame/internal/assets"
    "setcardgame/internal/graphics"
    "setcardgame/internal/ui"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
    ScreenWidth  = 800
    ScreenHeight = 600
)

type MenuScene struct {
    bg            *ebiten.Image
    buttons       []*ui.Button
    questionBtn   *ui.QuestionButton
    rulesWindow   *ui.RulesWindow
    roomKeyWindow *ui.RoomKeyWindow
    
    // Game state
    showingSplash bool
}

func NewMenuScene(bg, btnNormal, btnHover, squareNormal, squareHover *ebiten.Image, font *text.GoTextFaceSource) *MenuScene {
    // Load card images
    cardImages := make([]*ebiten.Image, 0)
    for _, cardData := range assets.GetAllCardData() {
        cardImg := graphics.LoadImageFromBytes(cardData)
        cardImages = append(cardImages, cardImg)
    }
    
    // Button dimensions
    buttonWidth := ScreenWidth / 5
    buttonHeight := ScreenHeight / 8
    buttonX := (ScreenWidth - buttonWidth) / 2
    
    // Square button dimensions (1/10 of screen size)
    squareButtonSize := ScreenWidth / 10
    squareButtonX := ScreenWidth - squareButtonSize - 10
    squareButtonY := 10
    
    // Position buttons vertically centered as a group
    totalButtonsHeight := buttonHeight*2 + 20
    startY := (ScreenHeight - totalButtonsHeight) / 2
    
    button1Y := startY
    button2Y := startY + buttonHeight + 20

    // Create scene instance first
    scene := &MenuScene{
        bg: bg,
    }

    // Rules window setup
    rulesWindow := ui.NewRulesWindow(160, 90, 480, 420, font)
    
    // Room key window setup (larger size)
    roomKeyWindow := ui.NewRoomKeyWindow(
        250, 150, 300, 250, // x, y, width, height (larger window)
        font, btnNormal, btnHover,
        func(roomKey string) {
            log.Printf("🚪 Joining room: %s", roomKey)
        },
        func() {
            log.Println("❌ Room join cancelled")
        },
    )
    

    // Create buttons
    button1 := ui.NewButton(
        buttonX, button1Y, buttonWidth, buttonHeight,
        btnNormal, btnHover, "Start", font,
        func() {
            log.Println("✅ Start button clicked!")
            roomKeyWindow.Open()
        },
    )

    button2 := ui.NewButton(
        buttonX, button2Y, buttonWidth, buttonHeight,
        btnNormal, btnHover, "Options", font,
        func() {
            log.Println("✅ Options button clicked!")
        },
    )

    questionButton := ui.NewQuestionButton(
        squareButtonX, squareButtonY, squareButtonSize, squareButtonSize,
        squareNormal, squareHover, font,
        func() {
            if rulesWindow.IsOpen || rulesWindow.Opening {
                rulesWindow.Close()
                log.Println("❌ Rules window closing")
            } else {
                rulesWindow.Open()
                log.Println("📖 Rules window opening")
            }
        },
    )

    // Set up the scene
    scene.buttons = []*ui.Button{button1, button2}
    scene.questionBtn = questionButton
    scene.rulesWindow = rulesWindow
    scene.roomKeyWindow = roomKeyWindow

    return scene
}

func (m *MenuScene) Update() error {
    // Update card splash first (it can override other inputs)
   

    // Normal menu updates
    for _, button := range m.buttons {
        button.Update()
    }
    m.questionBtn.Update()
    m.rulesWindow.Update()
    
    if m.roomKeyWindow != nil {
        m.roomKeyWindow.Update()
    }
    
    return nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
    // Draw background
    bgOpts := &ebiten.DrawImageOptions{}
    bgBounds := m.bg.Bounds()
    scaleX := float64(ScreenWidth) / float64(bgBounds.Dx())
    scaleY := float64(ScreenHeight) / float64(bgBounds.Dy())
    bgOpts.GeoM.Scale(scaleX, scaleY)
    screen.DrawImage(m.bg, bgOpts)
    

    
    // Normal menu drawing
    for _, button := range m.buttons {
        button.Draw(screen)
    }
    
    m.questionBtn.Draw(screen)
    m.rulesWindow.Draw(screen)
    
    if m.roomKeyWindow != nil {
        m.roomKeyWindow.Draw(screen)
    }
  }
func (m *MenuScene) Layout(_, _ int) (int, int) {
    return ScreenWidth, ScreenHeight
}

