package assets

import (
    _ "embed"
)

//go:embed background.png
var BackgroundPNG []byte

//go:embed UI_elements/Default@4x.png
var DefaultButtonPNG []byte

//go:embed UI_elements/Hover@4x.png
var HoverButtonPNG []byte

//go:embed UI_elements/Default@4x_square.png
var SquareDefaultPNG []byte

//go:embed UI_elements/Hover@4x_square.png
var SquareHoverPNG []byte

//go:embed FSEX300.ttf
var FontTTF []byte

