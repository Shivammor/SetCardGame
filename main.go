package main

import (
<<<<<<< HEAD
	"bytes"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X, Y          int
	Width, Height int
	NormalImage   *ebiten.Image
	HoverImage    *ebiten.Image
	Label         string
	FaceSource    *text.GoTextFaceSource
	OnClick       func()
}

func (b *Button) Update() {
	mx, my := ebiten.CursorPosition()
	if b.IsHovered(mx, my) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	// Choose image based on hover state
	mx, my := ebiten.CursorPosition()
	var currentImage *ebiten.Image
	if b.IsHovered(mx, my) {
		currentImage = b.HoverImage
	} else {
		currentImage = b.NormalImage
	}

	// Scale and draw button image to fit the button size
	opts := &ebiten.DrawImageOptions{}
	
	// Calculate scale to fit button dimensions
	imgBounds := currentImage.Bounds()
	scaleX := float64(b.Width) / float64(imgBounds.Dx())
	scaleY := float64(b.Height) / float64(imgBounds.Dy())
	
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(currentImage, opts)

	// prepare font face
	face := &text.GoTextFace{
		Source: b.FaceSource,
		Size:   20,
	}

	// measure text
	advance, _ := text.Measure(b.Label, face, 0)

	// Center text horizontally and vertically
	tx := float64(b.X) + (float64(b.Width)-advance)/2
	ty := float64(b.Y) + float64(b.Height)/2

	tOpts := &text.DrawOptions{}
	tOpts.GeoM.Translate(tx, ty)
	tOpts.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, b.Label, face, tOpts)
}

func (b *Button) IsHovered(mx, my int) bool {
	return mx >= b.X && mx < b.X+b.Width && my >= b.Y && my < b.Y+b.Height
}

type QuestionButton struct {
	X, Y          int
	Width, Height int
	NormalImage   *ebiten.Image
	HoverImage    *ebiten.Image
	FaceSource    *text.GoTextFaceSource
	OnClick       func()
}

func (q *QuestionButton) Update() {
	mx, my := ebiten.CursorPosition()
	if q.IsHovered(mx, my) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if q.OnClick != nil {
			q.OnClick()
		}
	}
}

func (q *QuestionButton) Draw(screen *ebiten.Image) {
	// Choose image based on hover state
	mx, my := ebiten.CursorPosition()
	var currentImage *ebiten.Image
	if q.IsHovered(mx, my) {
		currentImage = q.HoverImage
	} else {
		currentImage = q.NormalImage
	}

	// Scale and draw button image
	opts := &ebiten.DrawImageOptions{}
	imgBounds := currentImage.Bounds()
	scaleX := float64(q.Width) / float64(imgBounds.Dx())
	scaleY := float64(q.Height) / float64(imgBounds.Dy())
	
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(float64(q.X), float64(q.Y))
	screen.DrawImage(currentImage, opts)

	// Draw bold black question mark
	face := &text.GoTextFace{
		Source: q.FaceSource,
		Size:   32,
	}

	advance, _ := text.Measure("?", face, 0)
	tx := float64(q.X) + (float64(q.Width)-advance)/2
	ty := float64(q.Y) + float64(q.Height)/2

	tOpts := &text.DrawOptions{}
	tOpts.GeoM.Translate(tx, ty)
	tOpts.ColorScale.ScaleWithColor(color.Black)
	text.Draw(screen, "?", face, tOpts)
}

func (q *QuestionButton) IsHovered(mx, my int) bool {
	return mx >= q.X && mx < q.X+q.Width && my >= q.Y && my < q.Y+q.Height
}

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
	r.ContentHeight = float64(len(rules)) * lineHeight + margin*2
	
	// Calculate max scroll offset
	visibleHeight := height - margin*2
	if r.ContentHeight > visibleHeight {
		r.MaxScrollOffset = r.ContentHeight - visibleHeight
	} else {
		r.MaxScrollOffset = 0
	}

	// Create a clipping area for scrollable content
	//clipX := int(x + margin)
	//clipY := int(y + margin)
	//clipWidth := int(width - margin*2 - 15) // Leave space for scrollbar
	//clipHeight := int(height - margin*2)

	// Store original bounds for clipping
	//originalBounds := screen.Bounds()
	
	currentY := y + margin - r.ScrollOffset

	for i, rule := range rules {
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

type Game struct {
	bg            *ebiten.Image
	buttons       []*Button
	questionBtn   *QuestionButton
	rulesWindow   *RulesWindow
}

func (g *Game) Update() error {
	for _, button := range g.buttons {
		button.Update()
	}
	g.questionBtn.Update()
	g.rulesWindow.Update()
=======
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
>>>>>>> 45c61168719e604c59246808852360235b7c0f23
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
<<<<<<< HEAD
	// Scale background to cover entire screen
	bgOpts := &ebiten.DrawImageOptions{}
	bgBounds := g.bg.Bounds()
	scaleX := 800.0 / float64(bgBounds.Dx())
	scaleY := 600.0 / float64(bgBounds.Dy())
	bgOpts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.bg, bgOpts)
	
	for _, button := range g.buttons {
		button.Draw(screen)
	}
	
	g.questionBtn.Draw(screen)
	g.rulesWindow.Draw(screen)
}

func (g *Game) Layout(_, _ int) (int, int) { return 800, 600 }

func loadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func loadFont(path string) *text.GoTextFaceSource {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	
	src, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return src
}

func main() {
	bg := loadImage("./assets/background.png")
	btnNormalImg := loadImage("./assets/UI_elemets/Default@4x.png")
	btnHoverImg := loadImage("./assets/UI_elemets/Hover@4x.png")
	squareNormalImg := loadImage("./assets/UI_elemets/Default@4x_square.png")
	squareHoverImg := loadImage("./assets/UI_elemets/Hover@4x_square.png")
	fontSrc := loadFont("./FSEX300.ttf")

	// Button dimensions
	const screenWidth, screenHeight = 800, 600
	buttonWidth := screenWidth / 5
	buttonHeight := screenHeight / 8
	buttonX := (screenWidth - buttonWidth) / 2
	
	// Square button dimensions (1/10 of screen size)
	squareButtonSize := screenWidth / 10
	squareButtonX := screenWidth - squareButtonSize - 10
	squareButtonY := 10
	
	// Position buttons vertically centered as a group
	totalButtonsHeight := buttonHeight*2 + 20
	startY := (screenHeight - totalButtonsHeight) / 2
	
	button1Y := startY
	button2Y := startY + buttonHeight + 20

	// Rules window setup
	rulesWindow := &RulesWindow{
		X:         160,
		Y:         90,
		Width:     480,
		Height:    420,
		FaceSource: fontSrc,
		MaxFrames: 15,
	}

	button1 := &Button{
		X:           buttonX,
		Y:           button1Y,
		Width:       buttonWidth,
		Height:      buttonHeight,
		NormalImage: btnNormalImg,
		HoverImage:  btnHoverImg,
		Label:       "Start",
		FaceSource:  fontSrc,
		OnClick: func() {
			log.Println("✅ Start button clicked!")
		},
	}

	button2 := &Button{
		X:           buttonX,
		Y:           button2Y,
		Width:       buttonWidth,
		Height:      buttonHeight,
		NormalImage: btnNormalImg,
		HoverImage:  btnHoverImg,
		Label:       "Options",
		FaceSource:  fontSrc,
		OnClick: func() {
			log.Println("✅ Options button clicked!")
		},
	}

	questionButton := &QuestionButton{
		X:           squareButtonX,
		Y:           squareButtonY,
		Width:       squareButtonSize,
		Height:      squareButtonSize,
		NormalImage: squareNormalImg,
		HoverImage:  squareHoverImg,
		FaceSource:  fontSrc,
		OnClick: func() {
			if rulesWindow.IsOpen || rulesWindow.Opening {
				rulesWindow.Close()
				log.Println("❌ Rules window closing")
			} else {
				rulesWindow.Open()
				log.Println("📖 Rules window opening")
			}
		},
	}

	game := &Game{
		bg:          bg,
		buttons:     []*Button{button1, button2},
		questionBtn: questionButton,
		rulesWindow: rulesWindow,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine Button + Text")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

=======
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
>>>>>>> 45c61168719e604c59246808852360235b7c0f23
