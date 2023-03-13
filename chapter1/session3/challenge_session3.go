package main

import "fmt"

func main() {
	var countMap map[string]int = map[string]int{}
	word := ("Selamat pagi, Cikgu!")

	var isExist = func(key string, s string) bool {
		return key == s
	}

	var addValue = func(key string) {
		countMap[key] = countMap[key] + 1
	}

	var appendMap = func(s string) {
		countMap[s] = 1
	}

	var countChar = func(
		isExist func(key string, s string) bool,
		addValue func(key string),
		appendMap func(s string),
		s string) {
		for key := range countMap {
			if isExist(key, s) {
				addValue(key)
				goto BREAK
			}
		}

		appendMap(s)

	BREAK:
	}

	for _, s := range word {
		fmt.Println(string(s))
		countChar(isExist, addValue, appendMap, string(s))
	}

	fmt.Println(countMap)
}
