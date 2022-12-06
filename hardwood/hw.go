package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	woodsMap := make(map[string]int)
	//amounts := make([]float64, 10000, 10000)
	totalNumber := 0
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			wood := text

			if _, ok := woodsMap[wood]; ok {
				//amounts[woodsMap[wood]]++
				woodsMap[wood]++

			} else {
				woodsMap[wood] = 1
			}
			totalNumber++
		} else {
			var builder strings.Builder
			shortedKeys := shortMapKeys(woodsMap)

			for _, key := range shortedKeys {
				//builder.WriteString(fmt.Sprintf("%s %.6f \n", key, amounts[woodsMap[key]]/float64(totalNumber)*100))
				builder.WriteString(fmt.Sprintf("%s %.6f\n", key, float64(woodsMap[key])/float64(totalNumber)*100))
			}
			fmt.Println(builder.String())
			break
		}
	}
}

func shortMapKeys(woodsMap map[string]int) []string {
	keys := make([]string, 0, len(woodsMap))
	for key := range woodsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
