package dic

import (
	"bufio"
	"fmt"
	"os"
)

func Dictionary(charleng int) []string {

	var words []string

	f, err := os.Open("dic/wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		if len(scanner.Text()) == charleng {

			words = append(words, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return words
}
