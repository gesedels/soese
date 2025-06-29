///////////////////////////////////////////////////////////////////////////////////////
//              soese · stephen's over-engineered stack engine · v0.0.0              //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global state for our Forth machine
var (
	stack      []int
	dictionary map[string][]string
)

// Interpreter states
const (
	StateInterpret = iota
	StateCompile
)

// eval evaluates a single token. In Forth parlance, this is part of the "inner interpreter".
func eval(token string) error {
	// Is it a number?
	if num, err := strconv.Atoi(token); err == nil {
		stack = append(stack, num)
		return nil
	}

	// Is it a built-in word?
	switch token {
	case "+":
		if len(stack) < 2 {
			return fmt.Errorf("stack underflow: '+' requires two values")
		}
		a := stack[len(stack)-1]
		b := stack[len(stack)-2]
		stack = stack[:len(stack)-2] // Pop two
		stack = append(stack, b+a)   // Push sum
		return nil
	case ".":
		if len(stack) < 1 {
			return fmt.Errorf("stack underflow: '.' requires one value")
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop one
		fmt.Printf("%d ", val)
		return nil
	}

	// Is it a user-defined word?
	if def, ok := dictionary[token]; ok {
		// Execute the words in the definition
		for _, word := range def {
			if err := eval(word); err != nil {
				// Abort execution of this word on error
				return fmt.Errorf("in word '%s': %w", token, err)
			}
		}
		return nil
	}

	return fmt.Errorf("unknown word: %s", token)
}

func main() {
	// Initialize the machine state
	stack = make([]int, 0, 256)
	dictionary = make(map[string][]string)

	fmt.Println("///////////////////////////////////////////////////////////////////////////////////////")
	fmt.Println("//              soese · stephen's over-engineered stack engine · v0.0.0              //")
	fmt.Println("///////////////////////////////////////////////////////////////////////////////////////")
	fmt.Println("Minimal Forth-style interpreter. Enter 'bye' to exit.")

	scanner := bufio.NewScanner(os.Stdin)
	state := StateInterpret
	var newWordName string
	var newWordDef []string

	// The REPL (Read-Eval-Print Loop), or "outer interpreter"
	for {
		if state == StateInterpret {
			fmt.Printf("[ %v ]\n", stack)
			fmt.Print("ok> ")
		} else {
			// A more helpful prompt when defining a word
			fmt.Printf("... %s > ", newWordName)
		}

		if !scanner.Scan() {
			break // Exit on EOF (Ctrl+D)
		}

		line := scanner.Text()
		if line == "bye" {
			break
		}

		tokens := strings.Fields(line)
		tokenIdx := 0

		for tokenIdx < len(tokens) {
			token := tokens[tokenIdx]
			tokenIdx++

			if state == StateCompile {
				if token == ";" {
					// Finish definition
					dictionary[newWordName] = newWordDef
					state = StateInterpret
					newWordName = ""
					newWordDef = nil
				} else {
					// Add token to definition
					newWordDef = append(newWordDef, token)
				}
				continue
			}

			// In StateInterpret
			if token == ":" {
				if tokenIdx >= len(tokens) {
					fmt.Println("\nError: word name must follow ':' on the same line.")
					continue
				}
				state = StateCompile
				newWordName = tokens[tokenIdx]
				newWordDef = []string{}
				tokenIdx++ // Consume the name
				continue
			}

			if err := eval(token); err != nil {
				fmt.Printf("\nError: %v\n", err)
				// On error, clear the stack and stop processing the line.
				stack = stack[:0]
				break
			}
		}

		if state == StateInterpret {
			fmt.Println() // Newline after a successfully interpreted line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}

	fmt.Println("\nGoodbye!")
}
