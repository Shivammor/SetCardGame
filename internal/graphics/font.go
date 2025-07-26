package graphics

import (
    "bytes"
    "log"
    "os"

    "github.com/hajimehoshi/ebiten/v2/text/v2"
)

func LoadFontFromFile(path string) *text.GoTextFaceSource {
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

func LoadFontFromBytes(data []byte) *text.GoTextFaceSource {
    src, err := text.NewGoTextFaceSource(bytes.NewReader(data))
    if err != nil {
        log.Fatal(err)
    }
    return src
}

