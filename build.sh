GOOS=wasip1 GOARCH=wasm tinygo build -buildmode c-shared -o reactor.wasm main.go
GOOS=wasip1 GOARCH=wasm tinygo build -o command.wasm main.go