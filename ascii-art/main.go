package main

import (
	"fmt"
	"os"
	"strings"
)

func GetAscii(char int, file_read_bar string) string {
	ascii_word := char - 32
	file_read := strings.Split(file_read_bar, "\n\n")
	return file_read[ascii_word]
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Using: go run main.go <Your Word>")
		return
	}
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Printf("Error From Reading File: %v", err)
		return
	}

	content = content[1:]
	input_user := os.Args[1]
	// pam := 0
	printTxt(input_user, string(content))
}

func VirInput(s string) []string {
	var rStr []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			fmt.Println("Pleas Get a logic string <3 ")
			os.Exit(1)
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && s[i+1] == 'n' {
			rStr = append(rStr, word)
			i = i + 1
			word = ""
		} else {
			word = word + string(s[i])
		}
		if i == len(s)-1 {
			rStr = append(rStr, word)
		}
	}

	return rStr
}

func printH(s [][]string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < len(s); j++ {
			if i < len(s[j]) {
				fmt.Print(s[j][i])
			}

			// nStr[j][i] = string(sNew[x])
		}
		fmt.Println()
	}
}

func isEmpty(s []string)bool{
	for i := 0 ; i< len(s) ;i++{
		if s[i] !="" {
			return false
		}
	}
	return true
}

func printTxt(s, file string) {
	var nStr [][]string
	sNew := VirInput(s)
	for x := 0; x < len(sNew); x++ {
		s = string(sNew[x])
		if s == "" {
			if x != 0 || !isEmpty(sNew){
				fmt.Println()
			}
			continue
		}
		for i := 0; i < len(s); i++ {
			asciiStr := GetAscii(int(s[i]), file)
			if asciiStr != "" {
				lines := strings.Split(asciiStr, "\n")
				nStr = append(nStr, lines)
			} else {
				nStr = append(nStr, []string{})
			}
		}

		printH(nStr)
		nStr = nil
	}
}