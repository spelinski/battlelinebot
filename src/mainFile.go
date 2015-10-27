package main

import (
	//"github.com/stretchr/testify/assert"
	"bufio"
	"os"
	"parser"
)

func main() {
	mapReader := bufio.NewReader(os.Stdin)
	myParser := parser.Parser{}
	for {
		var input string
		isPrefix := true
		for isPrefix {
			var byteArray []byte
			byteArray, isPrefix, _ = mapReader.ReadLine()
			input = input + string(byteArray[:])
		}
		if len(input) > 0 {
			myParser.ParseString(input)
		}
	}
}
