

run:
	cd server && go run main.go

test-server:
	cd server && go test ./...

client:
	GOOS=js GOARCH=wasm go build -o asset/client.wasm github.com/kansi/go_wasm_todo/client


.PHONY: client
