package main

import "web-page-analyzer/internal/server"

// main is the application entry point.
// it is intentionally kept minimal and delegates
// all startup logic to the server package.
func main() {
	server.Start()
}
