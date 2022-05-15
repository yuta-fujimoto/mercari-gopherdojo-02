package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

  

func printStartScreen() {
	fmt.Print(" _               _                                          \n" +
		"| |_ _   _ _ __ (_)_ __   __ _    __ _  __ _ _ __ ___   ___ \n" +
		"| __| | | | '_ \\| | '_ \\ / _` |  / _` |/ _` | '_ ` _ \\ / _ \\\n" +
		"| |_| |_| | |_) | | | | | (_| | | (_| | (_| | | | | | |  __/\n" +
		" \\__|\\__, | .__/|_|_| |_|\\__, |  \\__, |\\__,_|_| |_| |_|\\___|\n" + 
		"     |___/|_|            |___/   |___/                      \n\n")
}

func initWords(words *[]string, fn string) {
	fp, err := os.Open(fn)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		*words = append(*words, scanner.Text())
	}
}

func doTypingGame(word string) bool {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(word)
	fmt.Print("-> ")

	scanner.Scan()

	return word == scanner.Text()
}

func main() {
	score := 0
	done := 0
	var words []string
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())

	initWords(&words, "data")
	printStartScreen()
	words_len := len(words)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nTime's up! Score: %d / %d\n", score, done)
			return
		default:
			if doTypingGame(words[rand.Intn(words_len)]) {
				score++
			}
			done++
		}
	}
}
