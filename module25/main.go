package main

import (
	"flag"
	"fmt"
)

func main() {
	var str, subStr string
	flag.StringVar(&str, "str", "", "set str")
	flag.StringVar(&subStr, "substr", "", "set substr")
	flag.Parse()

	fmt.Println(searchSubstr(str, subStr))
}

// searchSubstr - function for searching substring in string
func searchSubstr(str, subStr string) (result bool) {
	strRunes := []rune(str)
	subStrRunes := []rune(subStr)

	for i, stR := range strRunes {
		if result == true {
			break
		}
		if stR == subStrRunes[0] {
			if i+len(subStrRunes)-1 > len(strRunes) {
				break
			} else {
				for _, subR := range subStrRunes {
					if subR == strRunes[i] {
						result = true
					} else {
						result = false
						break
					}
					i++
				}
			}
		}
	}

	return
}
