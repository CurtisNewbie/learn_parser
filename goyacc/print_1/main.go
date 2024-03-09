package main

import (
	"bufio"
	"demo/parser"
	"os"
)

func main() {
	println("VM Running:\n")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inp := scanner.Text()
			if inp == "exit" {
				println("bye")
				return
			}
			parser.Parse(inp)
		}
	}
}
