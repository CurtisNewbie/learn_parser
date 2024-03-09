package main

import (
	"demo/parser"
)

func main() {
	s := `
		s = 'ok' // comments: s = '???'
		s = 'not ok'
		print  'reading s'
		print s
	`
	parser.Parse(s)
}
