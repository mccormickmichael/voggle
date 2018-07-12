package main

import (
	"fmt"
	"math/rand"
	"time"
	"tl/voggle/internal/board"
	"tl/voggle/internal/dict"
)

func main() {
	startTime := time.Now()
	
	b := board.RandomBoard(rand.NewSource(time.Now().Unix()))
	fmt.Println(b)
	
	d := dict.TreeDictFromPath("../../assets/dictionary.txt")
	
	c := board.Crawler{d}

	words := make(chan string)
	done := make(chan bool)

	foundWords := make(map[string]bool)
	waiters := 0
	// 5 goroutines
	// for row := 0; row < board.BoardSize; row++ {
	// 	go c.Crawl(b.Row(row), words, done)
	// 	waiters += 1
	// }
	// 25 goroutines, slightly faster
	for row := 0; row < board.BoardSize; row++ {
		for col := 0; col < board.BoardSize; col++ {
			go c.Crawl([]*board.Cell{b.At(row, col)}, words, done)
			waiters += 1
		}
	}

	for waiters > 0 {
		select {
		case w := <-words:
			if _, ok := foundWords[w]; !ok {
				fmt.Println(w)
			}
			foundWords[w] = true
		case <-done:
			waiters--
		}
	}
	fmt.Printf("%d words found in %s\n", len(foundWords), time.Since(startTime))
}
