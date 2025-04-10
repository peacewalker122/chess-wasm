GOOS=js GOARCH=wasm go build -o build/main.wasm go/main.go
cp $(go env GOROOT)/lib/wasm/wasm_exec.js build/
