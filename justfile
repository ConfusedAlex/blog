default:
    templ generate ./components/
    go run ./cmd/main.go serve
    xdg-open localhost:3000
