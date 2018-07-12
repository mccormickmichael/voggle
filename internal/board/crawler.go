package board

import (
//	"log"
	"tl/voggle/internal/dict"
)

const (
	MinWordLength = 4
)

type Crawler struct {
	Dict *dict.TreeDict
}

func (c *Crawler) Crawl(cells []*Cell, out chan<- string, done chan<- bool) {

	for _, cell := range cells {
		c.crawlInternal(cell, "", make([]*Cell, 10), out)
	}
	done <- true
}


func (c *Crawler) crawlInternal(cell *Cell, prefix string, visited []*Cell, out chan<- string) {
//	log.Printf("Entering %s", cell.String())
//	defer log.Printf("Exiting %s", cell.String())
	prefix += cell.Value
//	log.Printf("Checking %s", prefix)
	if !c.Dict.IsPrefix(prefix) {
//		log.Printf("%s does not start any word", prefix)
		return
	}
	if len(prefix) >= MinWordLength && c.Dict.IsWord(prefix) {
//		log.Printf("%s is a word!", prefix)
		out <- prefix
	}
	visited = append(visited, cell)
	for _, nc := range(cell.Neighbors) {
		visit := true
		for _, vc := range(visited) { // TODO: is there a better way to do this?
			if vc == nc {
//				log.Printf("%s already visited, skipping", nc.String())
				visit = false
			} 
		}
		if visit {
			c.crawlInternal(nc, prefix, visited, out)
		}
	}

}
