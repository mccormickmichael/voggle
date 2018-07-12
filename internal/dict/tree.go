package dict

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Node struct {
	value    byte
	parent   *Node
	endsWord bool
	children map[byte]*Node
}

type TreeDict struct {
	root  *Node
	count int
}

func EmptyTreeDict() *TreeDict {
	return &TreeDict{root: new(Node)}
}

func TreeDictFromPath(path string) *TreeDict {
	dict := EmptyTreeDict()
	df, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer df.Close()
	scanner := bufio.NewScanner(df)
	for scanner.Scan() {
		dict.Add(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return dict
}

func (dict *TreeDict) Words() int {
	return dict.count
}

func (dict *TreeDict) Add(word string) {
	if len(word) == 0 {
		return
	}
	dict.root.add(word)
	dict.count++
}

func (dict *TreeDict) IsWord(word string) bool {
	if len(word) == 0 {
		return false
	}
	node := dict.match(strings.ToLower(word))
	return node != nil && node.endsWord
}

func (dict *TreeDict) IsPrefix(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	node := dict.match(strings.ToLower(prefix))
	return node != nil
}

func (dict *TreeDict) match(word string) *Node {
	return dict.root.match(word)
}

func (dict *TreeDict) Breadth(queue chan *Node) {
	dict.root.breadth(queue)
}

func (dict *TreeDict) Depth(queue chan *Node) {
	dict.root.depth(queue)
	close(queue)
}

func (node *Node) String() string {
	return fmt.Sprintf("[%c word?:%t children:%d]", node.value, node.endsWord, len(node.children))
}

func (node *Node) add(word string) {
	if len(word) == 0 {
		return
	}
	if node.children == nil {
		node.children = make(map[byte]*Node)
	}
	first, remainder := split(word)

	childNode, ok := node.children[first]
	if !ok {
		childNode = &Node{value: first, parent: node}
		node.children[first] = childNode
	}
	if len(remainder) == 0 {
		childNode.endsWord = true
	}
	childNode.add(remainder)
}

func (node *Node) match(prefix string) *Node {
	if len(prefix) == 0 {
		return node
	}
	first, remainder := split(prefix)
	if child, ok := node.children[first]; ok {
		return child.match(remainder)
	}
	return nil
}

func (node *Node) IsWord() bool {
	return node.endsWord
}

func (node *Node) Reconstitute() string {
	if node.parent != nil {
		return node.parent.Reconstitute() + string(node.value)
	}
	return "" // special case for root node, which has no value
}

func (node *Node) Path() []string {
	if node.parent != nil {
		return append(node.parent.Path(), string(node.value))
	}
	return make([]string, 12) // the longest word in the dictionary we use has 12 letters
}

func (node *Node) breadth(queue chan *Node) {
	// WHAT TO PUSH ON TO THE CHANNEL?
	// get keys
	// sort them
	// push children
	// call breadth on each child
}

func (node *Node) depth(queue chan *Node) {

	if node.parent != nil { // special case, don't push the root node
		queue <- node
	}

	keys := node.childKeys()
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		n := node.children[k]
		n.depth(queue)
	}
}

func (node *Node) childKeys() []byte {
	keys := make([]byte, 0, len(node.children))
	for k := range node.children {
		keys = append(keys, k)
	}
	return keys
}

func split(suffix string) (byte, string) {
	return suffix[0], suffix[1:]
}
