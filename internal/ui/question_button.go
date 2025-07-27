package ui

import (
    "image/color"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/inpututil"
    "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type QuestionButton struct {
    X, Y          int
    Width, Height int
    NormalImage   *ebiten.Image
    HoverImage    *ebiten.Image
    FaceSource    *text.GoTextFaceSource
    OnClick       func()
}

func NewQuestionButton(x, y, width, height int, normalImg, hoverImg *ebiten.Image, font *text.GoTextFaceSource, onClick func()) *QuestionButton {
    return &QuestionButton{
        X:           x,
        Y:           y,
        Width:       width,
        Height:      height,
        NormalImage: normalImg,
        HoverImage:  hoverImg,
        FaceSource:  font,
        OnClick:     onClick,
    }
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
    ty := float64(q.Y) + float64(q.Height)/2 - 20

    tOpts := &text.DrawOptions{}
    tOpts.GeoM.Translate(tx, ty)
    tOpts.ColorScale.ScaleWithColor(color.Black)
    text.Draw(screen, "?", face, tOpts)
}

func (q *QuestionButton) IsHovered(mx, my int) bool {
    return mx >= q.X && mx < q.X+q.Width && my >= q.Y && my < q.Y+q.Height
}

