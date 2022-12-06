package main

import (
	"fmt"
	"strings"
)

const MaxNumberChar = 26

var alreadyVisited = [MaxNumberChar]bool{}

func main() {
	var m, n int
	_, _ = fmt.Scanf("%d%d\n", &m, &n)

	dictionaryMatrix := [MaxNumberChar][MaxNumberChar]bool{}
	for i := 0; i < m; i++ {
		var firstChar, secondChar rune
		_, _ = fmt.Scanf("%c %c\n", &firstChar, &secondChar)
		dictionaryMatrix[firstChar-'a'][secondChar-'a'] = true
	}

	var result strings.Builder
	for i := 0; i < n; i++ {
		var firstString, secondString string
		_, _ = fmt.Scanf("%s %s\n", &firstString, &secondString)

		if len(firstString) != len(secondString) {
			result.WriteString("no" + "\n")
		} else {
			isCorrect := true
			answer := "yes"
			for indexOfChar := 0; indexOfChar < len(firstString) && isCorrect; indexOfChar++ {
				if firstString[indexOfChar] != secondString[indexOfChar] {
					alreadyVisited[firstString[indexOfChar]-'a'] = true
					isCorrect, answer = check(dictionaryMatrix, int(firstString[indexOfChar]-'a'), int(secondString[indexOfChar]-'a'))
				}
			}
			result.WriteString(answer + "\n")
			alreadyVisited = [MaxNumberChar]bool{}
		}
	}
	fmt.Println(result.String())
}

func check(dictionaryMatrix [MaxNumberChar][MaxNumberChar]bool, firstChar int, secondChar int) (bool, string) {
	var isCorrect = false
	var answer = "no"
	for i := 0; i < MaxNumberChar && !isCorrect; i++ {
		if !alreadyVisited[i] && dictionaryMatrix[firstChar][i] {
			if i == secondChar {
				return true, "yes"
			} else {
				alreadyVisited[i] = true
				isCorrect, answer = check(dictionaryMatrix, i, secondChar)
			}
		}
	}
	return isCorrect, answer
}
