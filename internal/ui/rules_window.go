package ui

import (
    "image/color"
    "math"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
    "github.com/hajimehoshi/ebiten/v2/vector"
)

type RulesWindow struct {
    X, Y          int
    Width, Height int
    FaceSource    *text.GoTextFaceSource
    IsOpen        bool
    AnimFrame     int
    MaxFrames     int
    Opening       bool
    Closing       bool
    ScrollOffset  float64
    MaxScrollOffset float64
    ContentHeight float64
    rules         []string
}

func NewRulesWindow(x, y, width, height int, font *text.GoTextFaceSource) *RulesWindow {
    rules := []string{
        "Rules of Set Card Game",
        "",
        "1. Each player gets dealt 6 cards.",
        "2. The player who deals the cards gets 7 cards.",
        "3. The player who deals the cards goes first.",
        "4. The objective of the game is to minimize points.",
        "5. Each card has a point value.",
        "6. Cards numbers 2-10 are worth their number in points.",
        "7. Face Cards (Jack, Queen, King) are worth 10 points.",
        "8. Aces are worth 1 point.",
        "9. Players can create Set of cards that have zero points.",
        "10. A Set Can be made in the following ways:",
        "    a. All 3 or more cards are the same number.",
        "    b. All 3 or more cards are in a series of the same suit.",
        "11. Each player can at most have 6 cards at any time.",
        "12. Players can either draw a card from the deck or take",
        "    a card from the top of the discard pile.",
        "13. If any player has less or equal to 5 points, they can",
        "    close the round and can claim the win.",
        "14. If a player closes the round, all other players must",
        "    reveal their cards.",
        "15. If a player closes the round, they must show their",
        "    cards and calculate their points.",
        "16. Players can reduce their points further by making",
        "    sets with other players revealed cards.",
        "17. Over the course of rounds, players who have a total",
        "    of more than 52 points are eliminated.",
        "18. The game continues until only one player remains.",
    }

    return &RulesWindow{
        X:          x,
        Y:          y,
        Width:      width,
        Height:     height,
        FaceSource: font,
        MaxFrames:  15,
        rules:      rules,
    }
}

func (r *RulesWindow) Update() {
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

    // Handle scrolling when window is open
    if r.IsOpen {
        mx, my := ebiten.CursorPosition()
        // Check if mouse is over the window
        if mx >= r.X && mx <= r.X+r.Width && my >= r.Y && my <= r.Y+r.Height {
            _, dy := ebiten.Wheel()
            scrollSpeed := 30.0
            r.ScrollOffset -= dy * scrollSpeed
            
            // Clamp scroll offset
            if r.ScrollOffset < 0 {
                r.ScrollOffset = 0
            }
            if r.ScrollOffset > r.MaxScrollOffset {
                r.ScrollOffset = r.MaxScrollOffset
            }
        }
    }
}

func (r *RulesWindow) Open() {
    if !r.IsOpen && !r.Opening {
        r.Opening = true
        r.Closing = false
        r.ScrollOffset = 0 // Reset scroll when opening
    }
}

func (r *RulesWindow) Close() {
    if r.IsOpen && !r.Closing {
        r.Closing = true
        r.Opening = false
    }
}

func (r *RulesWindow) Draw(screen *ebiten.Image) {
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
    vector.DrawFilledRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), color.RGBA{240, 240, 240, 200}, false)
    
    // Draw border
    vector.StrokeRect(screen, float32(scaledX), float32(scaledY), float32(scaledWidth), float32(scaledHeight), 3, color.RGBA{50, 50, 50, 255}, false)

    // Only draw text if window is reasonably sized
    if scale > 0.3 {
        r.drawRules(screen, scaledX, scaledY, scaledWidth, scaledHeight, scale)
        
        // Draw scroll indicator if needed
        if r.MaxScrollOffset > 0 && r.IsOpen {
            r.drawScrollbar(screen, scaledX, scaledY, scaledWidth, scaledHeight)
        }
    }
}

func (r *RulesWindow) drawScrollbar(screen *ebiten.Image, x, y, width, height float64) {
    // Scrollbar dimensions
    scrollbarWidth := float32(8)
    scrollbarX := float32(x + width - 15)
    scrollbarY := float32(y + 10)
    scrollbarHeight := float32(height - 20)
    
    // Draw scrollbar track
    vector.DrawFilledRect(screen, scrollbarX, scrollbarY, scrollbarWidth, scrollbarHeight, color.RGBA{200, 200, 200, 150}, false)
    
    // Calculate thumb position and size
    thumbHeight := scrollbarHeight * float32(height) / float32(r.ContentHeight)
    if thumbHeight < 20 {
        thumbHeight = 20 // Minimum thumb size
    }
    
    thumbPosition := scrollbarY + (scrollbarHeight-thumbHeight)*float32(r.ScrollOffset/r.MaxScrollOffset)
    
    // Draw scrollbar thumb
    vector.DrawFilledRect(screen, scrollbarX, thumbPosition, scrollbarWidth, thumbHeight, color.RGBA{100, 100, 100, 200}, false)
}

func (r *RulesWindow) drawRules(screen *ebiten.Image, x, y, width, height, scale float64) {
    face := &text.GoTextFace{
        Source: r.FaceSource,
        Size:   14 * scale,
    }

    titleFace := &text.GoTextFace{
        Source: r.FaceSource,
        Size:   18 * scale,
    }

    margin := 20.0 * scale
    lineHeight := 20.0 * scale
    
    // Calculate total content height
    r.ContentHeight = float64(len(r.rules)) * lineHeight + margin*2
    
    // Calculate max scroll offset
    visibleHeight := height - margin*2
    if r.ContentHeight > visibleHeight {
        r.MaxScrollOffset = r.ContentHeight - visibleHeight
    } else {
        r.MaxScrollOffset = 0
    }

    currentY := y + margin - r.ScrollOffset

    for i, rule := range r.rules {
        // Skip rendering lines that are completely outside the visible area
        lineBottom := currentY + lineHeight
        if lineBottom < y+margin || currentY > y+height-margin {
            currentY += lineHeight
            continue
        }

        var currentFace *text.GoTextFace
        if i == 0 { // Title
            currentFace = titleFace
        } else {
            currentFace = face
        }

        // Only draw if within visible bounds
        if currentY >= y+margin-lineHeight && currentY <= y+height-margin {
            tOpts := &text.DrawOptions{}
            tOpts.GeoM.Translate(x+margin, currentY)
            
            if i == 0 { // Title in black
                tOpts.ColorScale.ScaleWithColor(color.Black)
            } else {
                tOpts.ColorScale.ScaleWithColor(color.RGBA{30, 30, 30, 255})
            }

            // Create a sub-image for clipping
            if currentY >= y+margin && currentY <= y+height-margin-lineHeight {
                text.Draw(screen, rule, currentFace, tOpts)
            }
        }
        currentY += lineHeight
    }
}

