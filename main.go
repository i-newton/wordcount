package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	sentence, err := readInput()
	if err != nil {
		fail(err)
	}
	words := strings.Fields(sentence)
	if err != nil {
		fail(err)
	}
	fmt.Println(len(words))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (sentence string, err error) {
	args := os.Args[1:]
	sentence = strings.Join(args, "")
	return sentence, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
