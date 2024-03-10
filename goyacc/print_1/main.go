package main

import (
	"bufio"
	"demo/parser"
	"flag"
	"os"
)

var (
	file *string
)

func init() {
	file = flag.String("file", "", "path to script file")
}

func main() {
	flag.Parse()
	if *file == "" {
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

	f, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	parser.Run(string(f))
}
