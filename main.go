package main

import (
	"fmt"
	"strings"
)

func main() {

	if ping() {
		for {
			inputs := strings.Fields(getInputCommand())

			switch command := inputs[0]; command {
			case "help":
				helperFunc()
			case "show":
				switch target := inputs[1]; target {
				case "collections":
					fmt.Println("Showing collections")
				}
				fmt.Println("Showing something...")
			default:
				fmt.Printf("Nothing selected")
			}
		}
	}
}
