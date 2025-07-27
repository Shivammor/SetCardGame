package ui

import (
    "image/color"
    "math"
    "unicode"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
)

type RoomKeyWindow struct {
    X, Y          int
    Width, Height int
    FaceSource    *text.GoTextFaceSource
    IsOpen        bool
    AnimFrame     int
    MaxFrames     int
    Opening       bool
    Closing       bool
    
    // Text input fields
    roomKey       string
    maxKeyLength  int
    cursorVisible bool
    cursorTimer   int
    
    // Enter button
    enterButton   *Button
    
    // Callbacks
    OnEnter       func(roomKey string)
    OnCancel      func()
}

func NewRoomKeyWindow(x, y, width, height int, font *text.GoTextFaceSource, 
    btnNormal, btnHover *ebiten.Image, onEnter func(string), onCancel func()) *RoomKeyWindow {
    
    // Create enter button (centered at bottom of window)
    buttonWidth := 100
    buttonHeight := 40
    buttonX := x + (width-buttonWidth)/2
    buttonY := y + height - 120
    
    enterButton := NewButton(
        buttonX, buttonY, buttonWidth, buttonHeight,
        btnNormal, btnHover, "Enter", font,
        nil,
      )
    
    return &RoomKeyWindow{
        X:            x,
        Y:            y,
        Width:        width,
        Height:       height,
        FaceSource:   font,
        MaxFrames:    15,
        maxKeyLength: 10,
        enterButton:  enterButton,
        OnEnter:      onEnter,
        OnCancel:     onCancel,
    }
}

func (r *RoomKeyWindow) Update() {
    // Handle animation
    if r.Opening && r.AnimFrame < r.MaxFrames {
        r.AnimFrame++
        if r.AnimFrame >= r.MaxFrames {
            r.Opening = false
            r.IsOpen = true
        }
    } else if r.Closing && r.AnimFrame > 0 {
        r.AnimFrame--
        if r.AnimFrame <= 0 {
            r.Closing = false
            r.IsOpen = false
        }
    }
    
    if !r.IsOpen {
        return
    }
    
    // Update cursor blink
    r.cursorTimer++
    if r.cursorTimer >= 30 { // Blink every 30 frames (0.5 seconds at 60 FPS)
        r.cursorVisible = !r.cursorVisible
        r.cursorTimer = 0
    }
    
    // Handle keyboard input
    r.handleKeyboardInput()
    
    // Update enter button with current room key
    r.enterButton.OnClick = func() {
        if len(r.roomKey) > 0 && r.OnEnter != nil {
            r.OnEnter(r.roomKey)
            r.Close()
        }
    }
    r.enterButton.Update()
    
    // Handle escape key to close
    if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
        if r.OnCancel != nil {
            r.OnCancel()
        }
        r.Close()
    }
    
    // Handle enter key
    if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(r.roomKey) > 0 {
        if r.OnEnter != nil {
            r.OnEnter(r.roomKey)
        }
        r.Close()
    }
}

func (r *RoomKeyWindow) handleKeyboardInput() {
    // Handle backspace
    if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
        if len(r.roomKey) > 0 {
            r.roomKey = r.roomKey[:len(r.roomKey)-1]
        }
        return
    }
    
    // Handle character input
    runes := ebiten.AppendInputChars(nil)
    for _, ru := range runes {
        if len(r.roomKey) >= r.maxKeyLength {
            break
        }
        
        // Only allow alphanumeric characters and some symbols
        if unicode.IsLetter(ru) || unicode.IsDigit(ru) || ru == '-' || ru == '_' {
            r.roomKey += string(ru)
        }
    }
}

func (r *RoomKeyWindow) Open() {
    if !r.IsOpen && !r.Opening {
        r.Opening = true
        r.Closing = false
        r.roomKey = "" // Reset room key when opening
        r.cursorVisible = true
        r.cursorTimer = 0
    }
}

func (r *RoomKeyWindow) Close() {
    if r.IsOpen && !r.Closing {
        r.Closing = true
        r.Opening = false
    }
}

func (r *RoomKeyWindow) Draw(screen *ebiten.Image) {
    if !r.IsOpen && !r.Opening && !r.Closing {
        return
    }

    // Calculate animation scale
    progress := float64(r.AnimFrame) / float64(r.MaxFrames)
    scale := progress

    // Apply easing for smoother animation
    scale = math.Sin(progress * math.Pi / 2)

    if scale <= 0 {
        return
    }

    // Calculate scaled dimensions
    scaledWidth := float64(r.Width) * scale
    scaledHeight := float64(r.Height) * scale
    scaledX := float64(r.X) + (float64(r.Width)-scaledWidth)/2
    scaledY := float64(r.Y) + (float64(r.Height)-scaledHeight)/2

    // Draw window background
    vector.DrawFilledRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), color.RGBA{240, 240, 240, 220}, false)
    
    // Draw border
    vector.StrokeRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), 3, color.RGBA{50, 50, 50, 255}, false)

    // Only draw content if window is reasonably sized
    if scale > 0.3 {
        r.drawContent(screen, scaledX, scaledY, scaledWidth, scaledHeight, scale)
    }
}

func (r *RoomKeyWindow) drawContent(screen *ebiten.Image, x, y, width, height, scale float64) {
    // Title
    titleFace := &text.GoTextFace{
        Source: r.FaceSource,
        Size:   18 * scale,
    }
    
    title := "Enter Room Key"
    titleAdvance, _ := text.Measure(title, titleFace, 0)
    titleX := x + (width-titleAdvance)/2
    titleY := y + 40*scale
    
    titleOpts := &text.DrawOptions{}
    titleOpts.GeoM.Translate(titleX, titleY)
    titleOpts.ColorScale.ScaleWithColor(color.Black)
    text.Draw(screen, title, titleFace, titleOpts)
    
    // Text input box
    inputBoxX := x + 20*scale
    inputBoxY := y + 80*scale
    inputBoxWidth := width - 40*scale
    inputBoxHeight := 40*scale
    
    // Draw input box background
    vector.DrawFilledRect(screen, float32(inputBoxX), float32(inputBoxY), float32(inputBoxWidth), float32(inputBoxHeight), color.RGBA{255, 255, 255, 255}, false)
    
    // Draw input box border
    vector.StrokeRect(screen, float32(inputBoxX), float32(inputBoxY), float32(inputBoxWidth), float32(inputBoxHeight), 2, color.RGBA{100, 100, 100, 255}, false)
    
    // Draw text inside input box
    textFace := &text.GoTextFace{
        Source: r.FaceSource,
        Size:   16 * scale,
    }
    
    displayText := r.roomKey
    if r.cursorVisible && r.IsOpen {
        displayText += "|"
    }
    
    if len(displayText) == 0 && !r.cursorVisible {
        // Show placeholder text
        placeholderOpts := &text.DrawOptions{}
        placeholderOpts.GeoM.Translate(inputBoxX+10*scale, inputBoxY+inputBoxHeight/2)
        placeholderOpts.ColorScale.ScaleWithColor(color.RGBA{150, 150, 150, 255})
        text.Draw(screen, "Room key...", textFace, placeholderOpts)
    } else {
        textOpts := &text.DrawOptions{}
        textOpts.GeoM.Translate(inputBoxX+10*scale, inputBoxY+inputBoxHeight/2)
        textOpts.ColorScale.ScaleWithColor(color.Black)
        text.Draw(screen, displayText, textFace, textOpts)
    }
    
    // Draw instructions
    instructionFace := &text.GoTextFace{
        Source: r.FaceSource,
        Size:   12 * scale,
    }
    
    instruction := "Press Enter to join or Escape to cancel"
    instrAdvance, _ := text.Measure(instruction, instructionFace, 0)
    instrX := x + (width-instrAdvance)/2
    instrY := y + height - 30*scale
    
    instrOpts := &text.DrawOptions{}
    instrOpts.GeoM.Translate(instrX, instrY)
    instrOpts.ColorScale.ScaleWithColor(color.RGBA{100, 100, 100, 255})
    text.Draw(screen, instruction, instructionFace, instrOpts)
    
    // Draw enter button
    if r.IsOpen && r.enterButton != nil {
        r.enterButton.Draw(screen)
    }
}

func (r *RoomKeyWindow) GetRoomKey() string {
    return r.roomKey
}

func (r *RoomKeyWindow) SetRoomKey(key string) {
    if len(key) <= r.maxKeyLength {
        r.roomKey = key
    }
}

