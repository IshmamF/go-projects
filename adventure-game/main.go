package main

import(
	"fmt"
	"go-projects/adventure-game/read"
)

const (
	filepath string = "gopher.json"
)

func main () {
	chapters := read.ReadJSON(filepath)
	fmt.Printf("%+v", chapters)
}