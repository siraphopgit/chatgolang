package main

import (
	"flag"

	"github.com/siraphopgit/chatgolang/chat"
)

var (
	port = flag.String("p", ":8000", "set port")
)

func init() {
	flag.Parse()
}

func main() {
	chat.Start(*port)
}
