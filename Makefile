# --------- Config ---------
GO_MAIN = go/main.go
WASM_OUTPUT = web/static/main.wasm
WASM_EXEC_JS = web/static/wasm_exec.js
WASM_EXEC_SRC = $(shell go env GOROOT)/lib/wasm/wasm_exec.js

# --------- Targets ---------

.PHONY: all build wasm svelte serve clean

all: build

## Build everything
build: wasm svelte

## Build Go → WASM
wasm:
	@echo "🛠️ Building Go WASM..."
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o $(WASM_OUTPUT) $(GO_MAIN)
	@echo "📦 Copying wasm_exec.js..."
	cp $(WASM_EXEC_SRC) $(WASM_EXEC_JS)

## Build Svelte frontend (uses Vite)
svelte:
	@echo "🎨 Building Svelte frontend..."
	cd web && npm install && npm run build

## Run dev server (SvelteKit/Vite)
serve:
	cd web && npm run dev

## Clean build outputs
clean:
	@echo "🧹 Cleaning WASM output..."
	rm -f $(WASM_OUTPUT) $(WASM_EXEC_JS)
	cd web && rm -rf dist .svelte-kit

