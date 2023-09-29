package app

import (
	"ASCII-ART-WEB/pkg/internals/check"
	"fmt"
	"os"
	"strings"
)

func Run(input, banner string) (string, int) {
	if !check.Ascii(input) {
		return "", 400
	}
	fmt.Println(input)
	elemsMap := make(map[rune][]string)
	switch banner {
	case "standard":
		banner = "./banners/standard.txt"
	case "shadow":
		banner = "./banners/shadow.txt"
	case "thinkertoy":
		banner = "./banners/thinkertoy.txt"
	}
	data, err := os.ReadFile(banner)
	if err != nil {
		return "", 500
	}

	sliceData := strings.Split(string(data), "\n")         // splits standart.txt into multiple substrings by enters
	input = strings.ReplaceAll(input, "\\n", string('\n')) // replace occurrences of the "\\n" with the newline character '\n'
	splittedArr := strings.Split(input, string('\n'))
	// termWidth := check.GetTerminalWidth()

	for i := ' '; i <= '~'; i++ {
		var strs []string
		for j := 0; j < 8; j++ {
			res := (int(i-' ') * 9) + j + 1
			strs = append(strs, sliceData[res])
		}
		elemsMap[i] = strs
	}

	res := ""
	if check.Valid(splittedArr) {
		for _, el := range splittedArr {
			if len(el) > 0 {
				line := getAsciiArtLine(el, elemsMap)
				// firstline := len(line) / 8
				// if firstline > termWidth {
				// 	fmt.Println("The output of your text does not fit in the terminal")
				// 	return "", 0
				// }
				res += line
			} else {
				res += string('\n')
			}
		}
	} else {
		for i := 0; i < len(splittedArr)-1; i++ { // handling empty input
			res = "\n" + res
		}
	}

	return res, 0
}

func getAsciiArtLine(str string, mapOfEl map[rune][]string) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, elem := range str {
			res += mapOfEl[elem][i]
		}
		res += "\n"
	}
	return res
}
