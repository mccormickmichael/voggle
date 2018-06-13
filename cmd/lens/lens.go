package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	histo := make(map[int]int, 10)
	var count, maxlen int

	df, err := os.Open("../dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()

	scanner := bufio.NewScanner(df)
	for scanner.Scan() {
		text := scanner.Text()
		count++
		wordlen := len(text)
		lencount := histo[wordlen]
		lencount++
		histo[wordlen] = lencount

		if wordlen > maxlen {
			maxlen = wordlen
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i <= maxlen; i++ {
		//		icount := histo[i]
		fmt.Printf("%2d %5d\n", i, histo[i])
	}
	fmt.Printf("%5d\n", count)
}
