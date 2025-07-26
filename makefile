.PHONY: build-desktop build-wasm serve clean

build-desktop:
	go build -o bin/game cmd/desktop/main.go

build-wasm:
	GOOS=js GOARCH=wasm go build -o web/game.wasm cmd/wasm/main.go

serve: build-wasm
	@echo "Game available at http://localhost:8080"
	cd web && python3 -m http.server 8080

clean:
	rm -rf bin/ web/game.wasm web/wasm_exec.js

run-desktop: build-desktop
	./bin/game

