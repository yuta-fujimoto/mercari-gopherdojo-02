package main

import (
	"bufio"
	"context"
	"fmt"
	"io"

	// "io"
	"os"
	"time"
	"typing_game/pickStr"
)

const (
	GREEN = "\033[32m"
	RED   = "\033[31m"
	RESET = "\033[0m"
)

const (
	deadline = 30 * time.Second
)

func printStartScreen() {
	fmt.Print(" _               _                                          \n" +
		"| |_ _   _ _ __ (_)_ __   __ _    __ _  __ _ _ __ ___   ___ \n" +
		"| __| | | | '_ \\| | '_ \\ / _` |  / _` |/ _` | '_ ` _ \\ / _ \\\n" +
		"| |_| |_| | |_) | | | | | (_| | | (_| | (_| | | | | | |  __/\n" +
		" \\__|\\__, | .__/|_|_| |_|\\__, |  \\__, |\\__,_|_| |_| |_|\\___|\n" +
		"     |___/|_|            |___/   |___/                      \n\n")
}

func runTypingGame(ctx context.Context, ch chan bool,
	scanner *bufio.Scanner) (bool, error) {
	want := pickStr.Pick()

	fmt.Println(want)
	fmt.Print("-> ")

	if scanner.Err() != nil {
		return false, scanner.Err()
	}

	go scan(ch, scanner)

	select {
	case got := <-ch:
		if got {
			return scanner.Text() == want, nil
		} else {
			return false, io.EOF
		}
	case <-ctx.Done():
		return false, nil
	}
}

func scan(in chan bool, scanner *bufio.Scanner) {
	if scanner.Scan() {
		in <- true
	} else {
		in <- false
	}
}

func main() {
	score := 0
	done := 0
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()

	channel := make(chan bool, 1)
	defer close(channel)
	scanner := bufio.NewScanner(os.Stdin)

	err := pickStr.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	printStartScreen()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nTime's up! Score: %d / %d\n", score, done)
			return
		default:
			res, err := runTypingGame(ctx, channel, scanner)
			if err != nil {
				fmt.Println("Sudden interruption...")
				return // EOF
			}
			if res {
				fmt.Println(GREEN, "Good job :)", RESET)
				score++
			} else {
				fmt.Println(RED, "Oops :(", RESET)
			}
			done++
		}
	}
}
