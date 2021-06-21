package main

import (
	"bufio"
	"cycloid-challenge-words/dic"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

///
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the fisrt word: ")
	first, _ := reader.ReadString('\n')
	first = strings.Replace(first, "\n", "", -1)

	fmt.Print("Enter the last word: ")
	last, _ := reader.ReadString('\n')
	last = strings.Replace(last, "\n", "", -1)

	startTime := time.Now()
	fmt.Println("START")
	// first := "cat"
	// last := "dog"

	if len(first) != len(last) {
		fmt.Println("Error: Incorrect Inputs")
		return
	}
	words := dic.Dictionary(len(first))

	NodePath := findChain(first, last, words)

	var path []string
	// fmt.Println(path)

	for {
		path = append(path, NodePath.word)

		NodePath = NodePath.parents[0]
		if NodePath.parents == nil {
			path = append(path, NodePath.word)
			break
		}
	}

	fmt.Println(reverseArray(path))
	fmt.Println("END")
	elapsed := time.Since(startTime)
	log.Printf("---> TIME %s", elapsed)

}

func findChain(first, last string, dicList []string) node {
	startNode := node{nil, first}
	var finalNode node
	checkIt := []node{startNode}
	var checked []string
	// var neigh []string

	// fmt.Println("+++", checkIt)

	for i := 0; i < len(checkIt); i++ {
		// fmt.Println("+++", i, len(checkIt))

		if checkIt[i].word == last {
			// fmt.Println(checkIt[i].parents)
			finalNode = checkIt[i]
		}
		queue := getSimilars(checkIt[i].word, dicList)

		for _, w := range queue {

			_, f := Find(checked, w)

			if !f {
				currentNode := node{[]node{checkIt[i]}, w}
				checkIt = append(checkIt, currentNode)

				checked = append(checked, w)
			}
		}

	}

	return finalNode
}

type node struct {
	parents []node
	word    string
}

func getSimilars(word string, universe []string) []string {
	var similars []string
	for _, w := range universe {
		if !diff(w, word) {
			similars = append(similars, w)
		}
	}
	return similars
}

func diff(word1, word2 string) bool {
	diffCount := 0

	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diffCount++
			if diffCount > 1 {
				return true
			}
		}
	}
	return false
}

// Find return true / false and the i
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
