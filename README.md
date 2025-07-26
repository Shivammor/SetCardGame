Set Card Game 🃏

A beautiful, multiplayer card game built with pure Go and Ebitengine, featuring seamless WASM deployment for web browsers.
✨ Features

    🎮 Pure Go Implementation - Built entirely in Go using the Ebitengine game engine

    🌐 WebAssembly Support - Runs natively in web browsers without plugins

    👥 Multiplayer Ready - Designed for multiplayer gameplay with real-time synchronization

    🎨 Beautiful UI - Smooth animations, hover effects, and responsive design

    📱 Cross-Platform - Works on desktop (Windows, macOS, Linux) and web

    🔧 Modular Architecture - Clean, maintainable codebase with separation of concerns

🎯 What Makes This Special
Pure Go Power

Unlike many games that mix multiple languages and frameworks, this project showcases the power of pure Go for game development:

    Single language for both client and server logic

    Strong typing and excellent tooling

    Easy cross-compilation and deployment

    Memory safety and garbage collection

WebAssembly Innovation

Experience the cutting-edge of web gaming:

    No JavaScript required - Pure Go compiled to WASM

    Near-native performance in the browser

    Asset embedding - All resources bundled into a single WASM file

    Instant loading - No external dependencies or downloads

Multiplayer Architecture

Built from the ground up for multiplayer experiences:

    Real-time game state synchronization

    Client-server architecture

    Optimistic updates for smooth gameplay

    Reconnection handling and game persistence

🎲 Game Rules

Set Card Game is a strategic card game where players aim to minimize their points:
Basic Rules

    Each player receives 6 cards (dealer gets 7)

    The dealer goes first

    Goal: Minimize your total points

Card Values

    Numbers 2-10: Face value points

    Face Cards (J, Q, K): 10 points each

    Aces: 1 point each

Creating Sets (0 Points)

Players can create sets with zero points:

    Same Number: 3+ cards of the same rank

    Straight Flush: 3+ consecutive cards of the same suit

Gameplay

    Draw from deck or take from discard pile

    Maximum 6 cards per player

    Close the round with ≤5 points to win

    Reduce points further using opponents' revealed cards

    Players with >52 total points are eliminated

    Last player standing wins!

🚀 Quick Start
Prerequisites

    Go 1.21 or higher

    Modern web browser (for WASM)

Running Locally (Desktop)

bash
git clone https://github.com/yourusername/set-card-game.git
cd set-card-game
make run-desktop

Building for Web (WASM)

bash
make build-wasm
make serve
# Navigate to http://localhost:8080

Development

bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build both targets
make build-desktop
make build-wasm

🏗️ Architecture
Project Structure

text
set-card-game/
├── cmd/
│   ├── desktop/        # Desktop entry point
│   └── wasm/          # WebAssembly entry point
├── internal/
│   ├── assets/        # Embedded game assets
│   ├── graphics/      # Image and font loading
│   ├── ui/           # Reusable UI components
│   ├── scene/        # Game scenes and states
│   ├── multiplayer/  # Networking and game logic
│   └── game/         # Core game mechanics
├── web/              # Web deployment files
└── assets/           # Source assets

Key Components
🎨 UI System

    Modular Components: Reusable buttons, windows, and controls

    Hover Effects: Dynamic visual feedback

    Smooth Animations: Eased transitions and scaling

    Responsive Design: Adapts to different screen sizes

🌐 Multiplayer Engine

    WebSocket Communication: Real-time bidirectional messaging

    Game State Management: Consistent state across all clients

    Event System: Decoupled game logic with event-driven architecture

    Reconnection Logic: Seamless recovery from network issues

🏃‍♂️ Performance

    Asset Embedding: Zero external file dependencies

    Efficient Rendering: Optimized draw calls and memory usage

    Hot Reload: Fast development iteration (desktop mode)

🌐 WASM Deployment

The game compiles to a single WASM file with all assets embedded:

bash
# Build produces a completely self-contained web app
make build-wasm

# Deploy the web/ directory to any static hosting:
# - Netlify
# - Vercel  
# - GitHub Pages
# - AWS S3
# - Your own server

Why WASM?

    Performance: Near-native speed in browsers

    Security: Sandboxed execution environment

    Compatibility: Works across all modern browsers

    Maintainability: Single codebase for desktop and web

🎮 Multiplayer Features
Current Implementation

    Real-time game state synchronization

    Player join/leave handling

    Turn-based gameplay

    Game room management

Planned Features

    Lobby system with room browser

    Spectator mode

    Tournament brackets

    Player statistics and rankings

    Chat system

    Mobile touch controls

🛠️ Development
Adding New Features

The modular architecture makes extending the game straightforward:

go
// Add new UI component
type NewComponent struct {
    // Component fields
}

// Implement Update() and Draw() methods
func (n *NewComponent) Update() { }
func (n *NewComponent) Draw(screen *ebiten.Image) { }

Asset Pipeline

Assets are automatically embedded at build time:

go
//go:embed new-asset.png
var NewAssetPNG []byte

Testing

bash
# Run all tests
go test ./...

# Test specific package
go test ./internal/game

# Run with coverage
go test -cover ./...

📦 Building & Distribution
Desktop Releases

bash
# Build for multiple platforms
GOOS=windows GOARCH=amd64 go build -o releases/game-windows.exe cmd/desktop/main.go
GOOS=darwin GOARCH=amd64 go build -o releases/game-macos cmd/desktop/main.go  
GOOS=linux GOARCH=amd64 go build -o releases/game-linux cmd/desktop/main.go

Web Deployment

bash
# Single command deployment
make build-wasm
# Upload web/ directory to your hosting provider


📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
🙏 Acknowledgments

    Ebitengine - Amazing 2D game engine for Go

    Go Team - For WebAssembly support and excellent tooling

    Contributors - Thanks to everyone who has contributed to this project!

🔗 Links

    🎮 Play Online - Try the game in your browser

    📖 Documentation - Detailed guides and API docs

    🐛 Report Issues - Found a bug? Let us know!

    💬 Discord Community - Join our gaming community

Built with ❤️ in Go | Powered by WebAssembly
