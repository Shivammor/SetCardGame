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

// Card backs and special cards
//go:embed playingcards/card-back1.png
var CardBack1PNG []byte

//go:embed playingcards/card-back2.png
var CardBack2PNG []byte

//go:embed playingcards/card-back3.png
var CardBack3PNG []byte

//go:embed playingcards/card-back4.png
var CardBack4PNG []byte

//go:embed playingcards/card-blank.png
var CardBlankPNG []byte

//go:embed playingcards/carddeck.png
var CardDeckPNG []byte

// Clubs (1-13)
//go:embed playingcards/card-clubs-1.png
var CardClubs1PNG []byte

//go:embed playingcards/card-clubs-2.png
var CardClubs2PNG []byte

//go:embed playingcards/card-clubs-3.png
var CardClubs3PNG []byte

//go:embed playingcards/card-clubs-4.png
var CardClubs4PNG []byte

//go:embed playingcards/card-clubs-5.png
var CardClubs5PNG []byte

//go:embed playingcards/card-clubs-6.png
var CardClubs6PNG []byte

//go:embed playingcards/card-clubs-7.png
var CardClubs7PNG []byte

//go:embed playingcards/card-clubs-8.png
var CardClubs8PNG []byte

//go:embed playingcards/card-clubs-9.png
var CardClubs9PNG []byte

//go:embed playingcards/card-clubs-10.png
var CardClubs10PNG []byte

//go:embed playingcards/card-clubs-11.png
var CardClubs11PNG []byte

//go:embed playingcards/card-clubs-12.png
var CardClubs12PNG []byte

//go:embed playingcards/card-clubs-13.png
var CardClubs13PNG []byte

// Diamonds (1-13)
//go:embed playingcards/card-diamonds-1.png
var CardDiamonds1PNG []byte

//go:embed playingcards/card-diamonds-2.png
var CardDiamonds2PNG []byte

//go:embed playingcards/card-diamonds-3.png
var CardDiamonds3PNG []byte

//go:embed playingcards/card-diamonds-4.png
var CardDiamonds4PNG []byte

//go:embed playingcards/card-diamonds-5.png
var CardDiamonds5PNG []byte

//go:embed playingcards/card-diamonds-6.png
var CardDiamonds6PNG []byte

//go:embed playingcards/card-diamonds-7.png
var CardDiamonds7PNG []byte

//go:embed playingcards/card-diamonds-8.png
var CardDiamonds8PNG []byte

//go:embed playingcards/card-diamonds-9.png
var CardDiamonds9PNG []byte

//go:embed playingcards/card-diamonds-10.png
var CardDiamonds10PNG []byte

//go:embed playingcards/card-diamonds-11.png
var CardDiamonds11PNG []byte

//go:embed playingcards/card-diamonds-12.png
var CardDiamonds12PNG []byte

//go:embed playingcards/card-diamonds-13.png
var CardDiamonds13PNG []byte

// Hearts (1-13)
//go:embed playingcards/card-hearts-1.png
var CardHearts1PNG []byte

//go:embed playingcards/card-hearts-2.png
var CardHearts2PNG []byte

//go:embed playingcards/card-hearts-3.png
var CardHearts3PNG []byte

//go:embed playingcards/card-hearts-4.png
var CardHearts4PNG []byte

//go:embed playingcards/card-hearts-5.png
var CardHearts5PNG []byte

//go:embed playingcards/card-hearts-6.png
var CardHearts6PNG []byte

//go:embed playingcards/card-hearts-7.png
var CardHearts7PNG []byte

//go:embed playingcards/card-hearts-8.png
var CardHearts8PNG []byte

//go:embed playingcards/card-hearts-9.png
var CardHearts9PNG []byte

//go:embed playingcards/card-hearts-10.png
var CardHearts10PNG []byte

//go:embed playingcards/card-hearts-11.png
var CardHearts11PNG []byte

//go:embed playingcards/card-hearts-12.png
var CardHearts12PNG []byte

//go:embed playingcards/card-hearts-13.png
var CardHearts13PNG []byte

// Spades (1-13)
//go:embed playingcards/card-spades-1.png
var CardSpades1PNG []byte

//go:embed playingcards/card-spades-2.png
var CardSpades2PNG []byte

//go:embed playingcards/card-spades-3.png
var CardSpades3PNG []byte

//go:embed playingcards/card-spades-4.png
var CardSpades4PNG []byte

//go:embed playingcards/card-spades-5.png
var CardSpades5PNG []byte

//go:embed playingcards/card-spades-6.png
var CardSpades6PNG []byte

//go:embed playingcards/card-spades-7.png
var CardSpades7PNG []byte

//go:embed playingcards/card-spades-8.png
var CardSpades8PNG []byte

//go:embed playingcards/card-spades-9.png
var CardSpades9PNG []byte

//go:embed playingcards/card-spades-10.png
var CardSpades10PNG []byte

//go:embed playingcards/card-spades-11.png
var CardSpades11PNG []byte

//go:embed playingcards/card-spades-12.png
var CardSpades12PNG []byte

//go:embed playingcards/card-spades-13.png
var CardSpades13PNG []byte

// Function to get all card data (full deck)
func GetAllCardData() [][]byte {
    return [][]byte{
        // Clubs
        CardClubs1PNG, CardClubs2PNG, CardClubs3PNG, CardClubs4PNG,
        CardClubs5PNG, CardClubs6PNG, CardClubs7PNG, CardClubs8PNG,
        CardClubs9PNG, CardClubs10PNG, CardClubs11PNG, CardClubs12PNG, CardClubs13PNG,
        
        // Diamonds
        CardDiamonds1PNG, CardDiamonds2PNG, CardDiamonds3PNG, CardDiamonds4PNG,
        CardDiamonds5PNG, CardDiamonds6PNG, CardDiamonds7PNG, CardDiamonds8PNG,
        CardDiamonds9PNG, CardDiamonds10PNG, CardDiamonds11PNG, CardDiamonds12PNG, CardDiamonds13PNG,
        
        // Hearts
        CardHearts1PNG, CardHearts2PNG, CardHearts3PNG, CardHearts4PNG,
        CardHearts5PNG, CardHearts6PNG, CardHearts7PNG, CardHearts8PNG,
        CardHearts9PNG, CardHearts10PNG, CardHearts11PNG, CardHearts12PNG, CardHearts13PNG,
        
        // Spades
        CardSpades1PNG, CardSpades2PNG, CardSpades3PNG, CardSpades4PNG,
        CardSpades5PNG, CardSpades6PNG, CardSpades7PNG, CardSpades8PNG,
        CardSpades9PNG, CardSpades10PNG, CardSpades11PNG, CardSpades12PNG, CardSpades13PNG,
    }
}

// Helper functions for specific card types
func GetCardBackData() [][]byte {
    return [][]byte{
        CardBack1PNG,
        CardBack2PNG,
        CardBack3PNG,
        CardBack4PNG,
    }
}

func GetAceCardData() [][]byte {
    return [][]byte{
        CardClubs1PNG, CardDiamonds1PNG, CardHearts1PNG, CardSpades1PNG,
    }
}

func GetFaceCardData() [][]byte {
    return [][]byte{
        // Jacks (11)
        CardClubs11PNG, CardDiamonds11PNG, CardHearts11PNG, CardSpades11PNG,
        // Queens (12)
        CardClubs12PNG, CardDiamonds12PNG, CardHearts12PNG, CardSpades12PNG,
        // Kings (13)
        CardClubs13PNG, CardDiamonds13PNG, CardHearts13PNG, CardSpades13PNG,
    }
}

func GetNumberCardData() [][]byte {
    return [][]byte{
        // 2s through 10s
        CardClubs2PNG, CardClubs3PNG, CardClubs4PNG, CardClubs5PNG,
        CardClubs6PNG, CardClubs7PNG, CardClubs8PNG, CardClubs9PNG, CardClubs10PNG,
        
        CardDiamonds2PNG, CardDiamonds3PNG, CardDiamonds4PNG, CardDiamonds5PNG,
        CardDiamonds6PNG, CardDiamonds7PNG, CardDiamonds8PNG, CardDiamonds9PNG, CardDiamonds10PNG,
        
        CardHearts2PNG, CardHearts3PNG, CardHearts4PNG, CardHearts5PNG,
        CardHearts6PNG, CardHearts7PNG, CardHearts8PNG, CardHearts9PNG, CardHearts10PNG,
        
        CardSpades2PNG, CardSpades3PNG, CardSpades4PNG, CardSpades5PNG,
        CardSpades6PNG, CardSpades7PNG, CardSpades8PNG, CardSpades9PNG, CardSpades10PNG,
    }
}

// Get special assets
func GetCardDeckPNG() []byte {
    return CardDeckPNG
}

func GetCardBlankPNG() []byte {
    return CardBlankPNG
}

