package ui

import (
    "image/color"
    "math"
    "setcardgame/internal/memory"
    "unicode"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
)

type PlayerSetupWindow struct {
    X, Y          int
    Width, Height int
    FaceSource    *text.GoTextFaceSource
    IsOpen        bool
    AnimFrame     int
    MaxFrames     int
    Opening       bool
    Closing       bool
    
    // Player data
    playerName      string
    maxNameLength   int
    selectedAvatar  int
    avatarImages    []*ebiten.Image
    
    // UI state
    cursorVisible   bool
    cursorTimer     int
    nameInputActive bool
    
    // Buttons
    confirmButton   *Button
    
    // Callbacks
    OnConfirm       func(name string, avatar int)
    OnCancel        func()
}

func NewPlayerSetupWindow(x, y, width, height int, font *text.GoTextFaceSource, 
    btnNormal, btnHover *ebiten.Image, avatarImages []*ebiten.Image,
    onConfirm func(string, int), onCancel func()) *PlayerSetupWindow {
    
    // Create confirm button (centered at bottom)
    buttonWidth := 120
    buttonHeight := 40
    buttonX := x + (width-buttonWidth)/2
    buttonY := y + height - 60
    
    confirmButton := NewButton(
        buttonX, buttonY, buttonWidth, buttonHeight,
        btnNormal, btnHover, "Confirm", font,
        nil, // Will be set in Update method
    )
    
    return &PlayerSetupWindow{
        X:              x,
        Y:              y,
        Width:          width,
        Height:         height,
        FaceSource:     font,
        MaxFrames:      15,
        maxNameLength:  15,
        selectedAvatar: 0,
        avatarImages:   avatarImages,
        confirmButton:  confirmButton,
        OnConfirm:      onConfirm,
        OnCancel:       onCancel,
        nameInputActive: true,
    }
}

func (p *PlayerSetupWindow) Update() {
    // Handle animation
    if p.Opening && p.AnimFrame < p.MaxFrames {
        p.AnimFrame++
        if p.AnimFrame >= p.MaxFrames {
            p.Opening = false
            p.IsOpen = true
        }
    } else if p.Closing && p.AnimFrame > 0 {
        p.AnimFrame--
        if p.AnimFrame <= 0 {
            p.Closing = false
            p.IsOpen = false
        }
    }
    
    if !p.IsOpen {
        return
    }
    
    // Update cursor blink
    p.cursorTimer++
    if p.cursorTimer >= 30 {
        p.cursorVisible = !p.cursorVisible
        p.cursorTimer = 0
    }
    
    // Handle keyboard input for name
    if p.nameInputActive {
        p.handleNameInput()
    }
    
    // Handle avatar selection with mouse
    p.handleAvatarSelection()
    
    // Update confirm button
    p.confirmButton.OnClick = func() {
        if len(p.playerName) > 0 && p.OnConfirm != nil {
            p.OnConfirm(p.playerName, p.selectedAvatar)
            p.Close()
        }
    }
    p.confirmButton.Update()
    
    // Handle escape key to cancel
    if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
        if p.OnCancel != nil {
            p.OnCancel()
        }
        p.Close()
    }
    
    // Handle enter key to confirm
    if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(p.playerName) > 0 {
        if p.OnConfirm != nil {
            p.OnConfirm(p.playerName, p.selectedAvatar)
        }
        p.Close()
    }
}

func (p *PlayerSetupWindow) handleNameInput() {
    // Handle backspace
    if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
        if len(p.playerName) > 0 {
            p.playerName = p.playerName[:len(p.playerName)-1]
        }
        return
    }
    
    // Handle character input
    runes := ebiten.AppendInputChars(nil)
    for _, ru := range runes {
        if len(p.playerName) >= p.maxNameLength {
            break
        }
        
        // Only allow letters, numbers, and spaces
        if unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == ' ' {
            p.playerName += string(ru)
        }
    }
}

func (p *PlayerSetupWindow) handleAvatarSelection() {
    mx, my := ebiten.CursorPosition()
    
    if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
        // Calculate avatar positions
        avatarSize := 80
        avatarSpacing := 100
        startX := p.X + (p.Width - (len(p.avatarImages)*avatarSpacing - (avatarSpacing-avatarSize))) / 2
        avatarY := p.Y + 180
        
        for i := 0; i < len(p.avatarImages); i++ {
            avatarX := startX + i*avatarSpacing
            
            // Check if click is within avatar circle
            centerX := avatarX + avatarSize/2
            centerY := avatarY + avatarSize/2
            distance := math.Sqrt(float64((mx-centerX)*(mx-centerX) + (my-centerY)*(my-centerY)))
            
            if distance <= float64(avatarSize/2) {
                p.selectedAvatar = i
                break
            }
        }
    }
}

func (p *PlayerSetupWindow) Open() {
    if !p.IsOpen && !p.Opening {
        p.Opening = true
        p.Closing = false
        
        // Load saved data
        savedData := memory.GetPlayerData()
        p.playerName = savedData.Name
        p.selectedAvatar = savedData.Avatar
        
        p.cursorVisible = true
        p.cursorTimer = 0
        p.nameInputActive = true
    }
}

func (p *PlayerSetupWindow) Close() {
    if p.IsOpen && !p.Closing {
        p.Closing = true
        p.Opening = false
    }
}

func (p *PlayerSetupWindow) Draw(screen *ebiten.Image) {
    if !p.IsOpen && !p.Opening && !p.Closing {
        return
    }

    // Calculate animation scale
    progress := float64(p.AnimFrame) / float64(p.MaxFrames)
    scale := math.Sin(progress * math.Pi / 2)

    if scale <= 0 {
        return
    }

    // Calculate scaled dimensions
    scaledWidth := float64(p.Width) * scale
    scaledHeight := float64(p.Height) * scale
    scaledX := float64(p.X) + (float64(p.Width)-scaledWidth)/2
    scaledY := float64(p.Y) + (float64(p.Height)-scaledHeight)/2

    // Draw window background
    vector.DrawFilledRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), 
        color.RGBA{240, 240, 240, 230}, false)
    
    // Draw border
    vector.StrokeRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), 
        3, color.RGBA{50, 50, 50, 255}, false)

    // Only draw content if window is reasonably sized
    if scale > 0.3 {
        p.drawContent(screen, scaledX, scaledY, scaledWidth, scaledHeight, scale)
    }
}

func (p *PlayerSetupWindow) drawContent(screen *ebiten.Image, x, y, width, height, scale float64) {
    // Title
    titleFace := &text.GoTextFace{
        Source: p.FaceSource,
        Size:   20 * scale,
    }
    
    title := "Setup Your Player"
    titleAdvance, _ := text.Measure(title, titleFace, 0)
    titleX := x + (width-titleAdvance)/2
    titleY := y + 30*scale
    
    titleOpts := &text.DrawOptions{}
    titleOpts.GeoM.Translate(titleX, titleY)
    titleOpts.ColorScale.ScaleWithColor(color.Black)
    text.Draw(screen, title, titleFace, titleOpts)
    
    // Name input label
    labelFace := &text.GoTextFace{
        Source: p.FaceSource,
        Size:   14 * scale,
    }
    
    nameLabel := "Enter Your Name:"
    labelOpts := &text.DrawOptions{}
    labelOpts.GeoM.Translate(x+20*scale, y+70*scale)
    labelOpts.ColorScale.ScaleWithColor(color.Black)
    text.Draw(screen, nameLabel, labelFace, labelOpts)
    
    // Name input box
    inputBoxX := x + 20*scale
    inputBoxY := y + 90*scale
    inputBoxWidth := width - 40*scale
    inputBoxHeight := 35*scale
    
    // Draw input box background
    vector.DrawFilledRect(screen, float32(inputBoxX), float32(inputBoxY), float32(inputBoxWidth), float32(inputBoxHeight), 
        color.RGBA{255, 255, 255, 255}, false)
    
    // Draw input box border
    borderColor := color.RGBA{100, 100, 100, 255}
    if p.nameInputActive {
        borderColor = color.RGBA{70, 130, 180, 255} // Blue when active
    }
    vector.StrokeRect(screen, float32(inputBoxX), float32(inputBoxY), float32(inputBoxWidth), float32(inputBoxHeight), 
        2, borderColor, false)
    
    // Draw name text
    textFace := &text.GoTextFace{
        Source: p.FaceSource,
        Size:   16 * scale,
    }
    
    displayText := p.playerName
    if p.cursorVisible && p.IsOpen && p.nameInputActive {
        displayText += "|"
    }
    
    if len(displayText) == 0 && !p.cursorVisible {
        // Show placeholder text
        placeholderOpts := &text.DrawOptions{}
        placeholderOpts.GeoM.Translate(inputBoxX+8*scale, inputBoxY+inputBoxHeight/2+2*scale)
        placeholderOpts.ColorScale.ScaleWithColor(color.RGBA{150, 150, 150, 255})
        text.Draw(screen, "Your name...", textFace, placeholderOpts)
    } else {
        textOpts := &text.DrawOptions{}
        textOpts.GeoM.Translate(inputBoxX+8*scale, inputBoxY+inputBoxHeight/2+2*scale)
        textOpts.ColorScale.ScaleWithColor(color.Black)
        text.Draw(screen, displayText, textFace, textOpts)
    }
    
    // Avatar selection label
    avatarLabel := "Choose Your Avatar:"
    avatarLabelOpts := &text.DrawOptions{}
    avatarLabelOpts.GeoM.Translate(x+20*scale, y+140*scale)
    avatarLabelOpts.ColorScale.ScaleWithColor(color.Black)
    text.Draw(screen, avatarLabel, labelFace, avatarLabelOpts)
    
    // Draw avatars
    p.drawAvatars(screen, x, y, width, scale)
    
    // Draw confirm button
    if p.IsOpen {
        p.confirmButton.Draw(screen)
    }
    
    // Instructions
    instructionFace := &text.GoTextFace{
        Source: p.FaceSource,
        Size:   12 * scale,
    }
    
    instruction := "Press Enter to confirm or Escape to cancel"
    instrAdvance, _ := text.Measure(instruction, instructionFace, 0)
    instrX := x + (width-instrAdvance)/2
    instrY := y + height - 20*scale
    
    instrOpts := &text.DrawOptions{}
    instrOpts.GeoM.Translate(instrX, instrY)
    instrOpts.ColorScale.ScaleWithColor(color.RGBA{100, 100, 100, 255})
    text.Draw(screen, instruction, instructionFace, instrOpts)
}

func (p *PlayerSetupWindow) drawAvatars(screen *ebiten.Image, x, y, width, scale float64) {
    if len(p.avatarImages) == 0 {
        return
    }
    
    avatarSize := int(80 * scale)
    avatarSpacing := int(100 * scale)
    startX := int(x) + (int(width) - (len(p.avatarImages)*avatarSpacing - (avatarSpacing-avatarSize))) / 2
    avatarY := int(y) + int(160*scale)
    
    for i, avatarImg := range p.avatarImages {
        if avatarImg == nil {
            continue
        }
        
        avatarX := startX + i*avatarSpacing
        centerX := avatarX + avatarSize/2
        centerY := avatarY + avatarSize/2
        
        // Draw selection circle background
        if i == p.selectedAvatar {
            // Selected avatar - draw blue background circle
            vector.DrawFilledCircle(screen, float32(centerX), float32(centerY), float32(avatarSize/2+4), 
                color.RGBA{70, 130, 180, 255}, false)
        } else {
            // Unselected avatar - draw gray background circle
            vector.DrawFilledCircle(screen, float32(centerX), float32(centerY), float32(avatarSize/2+2), 
                color.RGBA{200, 200, 200, 255}, false)
        }
        
        // Draw avatar image in circle
        opts := &ebiten.DrawImageOptions{}
        
        // Calculate scale to fit in circle
        bounds := avatarImg.Bounds()
        imgScale := float64(avatarSize) / float64(bounds.Dx())
        if float64(avatarSize) / float64(bounds.Dy()) < imgScale {
            imgScale = float64(avatarSize) / float64(bounds.Dy())
        }
        
        opts.GeoM.Scale(imgScale, imgScale)
        opts.GeoM.Translate(float64(avatarX), float64(avatarY))
        
        screen.DrawImage(avatarImg, opts)
        
        // Draw circle border
        borderColor := color.RGBA{150, 150, 150, 255}
        if i == p.selectedAvatar {
            borderColor = color.RGBA{70, 130, 180, 255}
        }
        vector.StrokeCircle(screen, float32(centerX), float32(centerY), float32(avatarSize/2), 3, borderColor, false)
    }
}

func (p *PlayerSetupWindow) GetPlayerName() string {
    return p.playerName
}

func (p *PlayerSetupWindow) GetSelectedAvatar() int {
    return p.selectedAvatar
}

