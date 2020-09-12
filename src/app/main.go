package main

import (
	"app/portrait"
	"fmt"
)

func main() {
	fmt.Println("starting server at port 8100")
	portrait.NewPortraitApp()
}
