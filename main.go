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
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Printf("%s@%s > ", os.Getenv("USER"), hostname)

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func execInput(input string) error {
	// remove new line
	input = strings.TrimSuffix(input, "\n")
	// split the input
	args := strings.Split(input, " ")
	// gopher builtin commands
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir((args[1]))
	case "exit":
		os.Exit(0)
	}
	// create command
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
