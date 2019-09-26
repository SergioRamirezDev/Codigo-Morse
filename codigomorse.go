package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	var triminput string
	buf := bytes.Buffer{}
	numbers := [][]string{[]string{"1"}, []string{"2", "a", "b", "c"}, []string{"3", "d", "e", "f"}, []string{"4", "g", "h", "i"}, []string{"5", "j", "k", "l"}, []string{"6", "m", "n", "o"}, []string{"7", "p", "q", "r", "s"}, []string{"8", "t", "u", "v"}, []string{"9", "x", "y", "z"}, []string{"0", " "}}
	fmt.Print("Menu: \nType 1 or 2 for the following options\n(1) Numbers to letters\n(2) Letters to Numbers\n(3) Exit\n")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Print("Input: ")
		fmt.Scanln(&input)
		triminput = strings.TrimSpace(input)
		if triminput != "" {
			arr := splitGroupsOfNumbers(triminput)
			for index := 0; index < len(arr); index++ {
				buf.WriteString(convertCodeToLetter(arr[index], numbers))
			}
		}
	case "2":
		fmt.Print("Input: ")
		fmt.Scanln(&input)
		triminput = strings.TrimSpace(input)
		if triminput != "" {
			arr := strings.Split(triminput, "")
			for index := 0; index < len(arr); index++ {
				buf.WriteString(convertLetterToCode(arr[index], numbers))
			}
		}
		break
	case "3":
		os.Exit(3)
		break
	}
	result := buf.String()
	fmt.Println(result)
}

func convertLetterToCode(letter string, numbers [][]string) string {
	result := []string{}
	for x := 0; x < len(numbers); x++ {
		for y := 0; y < len(numbers[x]); y++ {
			if numbers[x][y] == letter {
				for c := 0; c < y+1; c++ {
					result = append(result, strconv.Itoa(x+1))
				}
				return strings.Join(result, "")
			}
		}
	}
	return ""
}

func convertCodeToLetter(code string, numbers [][]string) string {
	for index := 0; index < len(numbers); index++ {
		if numbers[index][0] == code[0:1] {
			if len(numbers[index]) >= len(code) {
				return string(numbers[index][len(code)-1])
			}
		}
	}
	return ""
}

func splitGroupsOfNumbers(triminput string) []string {
	stmcompare := []int{0}
	arrinput := strings.Split(triminput, "")
	result := []string{}
	for index := 0; index < len(arrinput); index++ {
		if index+1 < len(arrinput) && arrinput[index] != arrinput[index+1] {
			stmcompare = append(stmcompare, index+1)
		}
	}
	for index := 0; index < len(stmcompare); index++ {
		if index == 0 {
			result = append(result, triminput[:stmcompare[index+1]])
		} else if index+1 == len(stmcompare) {
			result = append(result, triminput[stmcompare[index]:])
		} else {
			result = append(result, triminput[stmcompare[index]:stmcompare[index+1]])
		}
	}
	return result
}
