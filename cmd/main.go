package main

import (
	http_ "english_bot_admin/internal/httpServer"
)

func main() {
	server := http_.NewServer()
	server.Run()
}
