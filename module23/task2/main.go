package main

import (
	"fmt"
	"strings"
)

// parseTest - function for parsing the last word in a sentence into runes
func parseTest(sentences [4]string, chars [5]rune) ([4][5]int, error) {
	var answerArray [4][5]int
	if len(sentences) == 0 || len(chars) == 0 {
		return answerArray, fmt.Errorf("the array is empty")
	}

	for i, s := range sentences {
		s = strings.ToUpper(s)
		words := strings.Fields(s)
		lastWord := words[len(words)-1]
		for j, ch := range chars {
			index := -1
			if strings.Contains(lastWord, string(ch)) {
				index = strings.Index(lastWord, string(ch))
			}
			answerArray[i][j] = index
		}
	}

	return answerArray, nil
}

func main() {
	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}
	answerArray, err := parseTest(sentences, chars)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, s := range sentences {
		fmt.Println("Sentence:", s)
		for j, ch := range chars {
			if answerArray[i][j] >= 0 {
				fmt.Printf("'%c' position %d\n", ch, answerArray[i][j])
			}
		}
	}
}
