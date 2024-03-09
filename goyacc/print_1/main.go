package main

import (
	"bufio"
	"demo/parser"
	"os"
)

func main() {
	// s := `
	// 	s = 'ok' // comments: s = '???'
	// 	s = 'not ok'
	// 	print  'reading s'
	// 	print s
	// `

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
