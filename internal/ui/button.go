package ui

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
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

func NewButton(x, y, width, height int, normalImg, hoverImg *ebiten.Image, label string, font *text.GoTextFaceSource, onClick func()) *Button {
    return &Button{
        X:           x,
        Y:           y,
        Width:       width,
        Height:      height,
        NormalImage: normalImg,
        HoverImage:  hoverImg,
        Label:       label,
        FaceSource:  font,
        OnClick:     onClick,
    }
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

