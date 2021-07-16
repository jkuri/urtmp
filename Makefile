all: build

build:
	@make build_ui
	@make statik
	@CGO_ENABLED=0 go build -o ./build/urtmp ./cmd/urtmp/main.go

build_ui:
	@if [ ! -d "web/urtmp/public/build" ]; then cd web/urtmp && npm run build; fi

statik:
	@if [ ! -r "internal/ui/statik.go" ]; then statik -dest ./internal -p ui -src ./web/urtmp/public; fi

clean:
	@rm -rf ./build ./web/urtmp/public/build ./internal/ui

install:
	@go install github.com/jkuri/statik/...@latest

.PHONY: build build_ui statik clean install
