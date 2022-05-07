package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func initWords(words *[]string, fn string) {
	fp, err := os.Open(fn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		*words = append(*words, scanner.Text())
	}
}

func doTypingGame(r io.Reader, w io.Writer, word string) bool {
	scanner := bufio.NewScanner(r)

	w.Write([]byte(word + "\n"))
	w.Write([]byte("-> "))
	scanner.Scan()
	return word == scanner.Text()
}

func main() {
	score := 0
	done := 0
	var isCorrect bool
	var words []string
	var words_len int
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())

	initWords(&words, "data")
	words_len = len(words)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Time's up! Score: %d / %d\n", score, done)
			return
		default:
			isCorrect = doTypingGame(os.Stdin, os.Stdout, words[rand.Intn(words_len)])
			if isCorrect {
				score++
			}
			done++
		}
	}
}
