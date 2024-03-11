package main

import (
	"bufio"
	"demo/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		println(">>> Entered interactive mode:")
		print(">>> ")
		for {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				inp := scanner.Text()
				if inp == "exit" {
					println("bye")
					return
				}
				parser.Run(inp)
				print(">>> ")
			}
		}
	}

	fmt.Printf(">>> Reading file %v\n", os.Args[1])
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	parser.Run(string(f))
}
