package main

import "demo/parser"

func main() {
	s := "s = 'ok'"
	parser.Parse(s)
	println("parse done")
}
