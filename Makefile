all: build

build:
	@make build_ui
	@make statik
	@CGO_ENABLED=0 go build -o ./build/urtmp ./cmd/urtmp/main.go

build_ui:
	@if [ ! -d "web/urtmp/dist/urtmp" ]; then cd web/urtmp && npm run build; fi

statik:
	@if [ ! -r "internal/ui/statik.go" ]; then statik -dest ./internal -p ui -src ./web/urtmp/dist/urtmp; fi

clean:
	@rm -rf ./build ./web/urtmp/dist/urtmp ./internal/ui

install:
	@go install github.com/jkuri/statik/...@latest
	@cd web/urtmp && npm install

docker:
	@docker build -t jkuri/urtmp .

.PHONY: build build_ui statik clean install docker
