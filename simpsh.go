// ALL HAIL SIMPSH
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("simpsh> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// remove the '\n'
	input = strings.TrimSuffix(input, "\n")
	// split the input to separate the command and arguments
	args := strings.Split(input, " ")

	// check for built-in commands
	switch args[0] {
	case "cd":
		// change dir to home dir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("Path required")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		// stop further processing
		return nil
	case "exit":
		os.Exit(0)
	case "quit":
		os.Exit(0)
	case "::version":
		fmt.Println("YAAY SIMPSH 1.0")
		return nil
	}

	// execute the cmd and args
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
