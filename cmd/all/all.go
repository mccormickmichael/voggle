package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"tl/voggle/internal/dict"
)

func main() {

	start := time.Now()

	td := dict.EmptyTreeDict()

	df, err := os.Open("../../assets/dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()

	scanner := bufio.NewScanner(df)
	for scanner.Scan() {
		td.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d words in dictionary!\n", td.Words())

	fmt.Printf("loading took %s\n", time.Since(start))
}
