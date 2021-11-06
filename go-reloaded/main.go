package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func start(sb string, su []string) bool {
	for _, si := range su {
		if strings.HasPrefix(sb, si) {
			return true
		}
	}
	return false
}
func toString(nbr int64) string {
	ini := int(nbr)
	string_int := strconv.Itoa(ini)
	return string_int
}
func toInt(sbr string) int {
	n1, err := strconv.Atoi(sbr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return n1
}
func HexToDec(array []string) []string {
	for pos, cb := range array {
		if cb == "(hex)" {
			value1, err := strconv.ParseInt(array[pos-1], 16, 64)
			if err != nil {
				fmt.Printf("Conversion failed: %s\n", err)
			}
			array[pos-1] = toString(value1)
			array[pos] = ""
		}
	}
	array = TrimEmpties(array)
	return array
}
func BinToDec(array []string) []string {
	for po, cc := range array {
		if cc == "(bin)" {
			value, err := strconv.ParseInt(array[po-1], 2, 64)
			if err != nil {
				fmt.Printf("Conversion failed: %s\n", err)
			}
			array[po-1] = toString(value)
			array[po] = ""
		}
	}
	array = TrimEmpties(array)
	return array
}
func ToUp(array []string) []string {
	for po, word1 := range array {
		if word1 == "(up)" {
			array[po-1] = strings.ToUpper(array[po-1])
			array[po] = ""
		}
		if word1 == "(up," {
			// parameter without )
			parameter := array[po+1][:len(array[po+1])-1]
			number := toInt(parameter)
			if number > len(array) {
				fmt.Println("Error with length")
				os.Exit(1)
			}
			for i := po - number; i < po; i++ {
				array[i] = strings.ToUpper(array[i])
			}
			array[po] = ""
			array[po+1] = ""
		}
	}
	array = TrimEmpties(array)
	return array
}
func ToLow(array []string) []string {
	for pos, word := range array {
		if word == "(low)" {
			array[pos-1] = strings.ToLower(array[pos-1])
			array[pos] = ""
		}
		if word == "(low," {
			// parameter without )
			parameter := array[pos+1][:len(array[pos+1])-1]
			number := toInt(parameter)
			if number > len(array) {
				fmt.Println("Error with length")
				os.Exit(1)
			}
			for i := pos - number; i < pos; i++ {
				array[i] = strings.ToLower(array[i])
			}
			array[pos] = ""
			array[pos+1] = ""
		}
	}
	array = TrimEmpties(array)
	return array
}
func Capitalize(array []string) []string {
	for pos, word := range array {
		if word == "(cap)" {
			array[pos-1] = strings.Title(array[pos-1])
			array[pos] = ""
		}
		if word == "(cap," {
			// parameter without )
			parameter := array[pos+1][:len(array[pos+1])-1]
			number := toInt(parameter)
			if number > len(array) {
				fmt.Println("Error with length")
				os.Exit(1)
			}
			for i := pos - number; i < pos; i++ {
				array[i] = strings.Title(array[i])
			}
			array[pos] = ""
			array[pos+1] = ""
		}
	}
	array = TrimEmpties(array)
	return array
}
func TrimEmpties(array []string) []string {
	var new_arr []string
	for _, word := range array {
		if word != "" {
			new_arr = append(new_arr, word)
		}
	}
	array = new_arr
	return array
}
func AddSpace(s string) string {
	new_word := ""
	new_word += string(s[0])
	new_word += string(rune(32))
	// new_word += 32
	for i := 1; i < len(s); i++ {
		new_word += string(s[i])
	}
	s = string(new_word)
	return s
}
func checkAlpha(s string) bool {
	last_char := s[len(s)-1]
	if (last_char > 'a' && last_char < 'z') || (last_char > 'A' && last_char < 'Z') {
		return true
	}
	return false
}
func Ponctuation(array []string) []string {
	puncs := []string{",", ".", ";", ":", "!", "?", "!?", "..."}
	for pos, word := range array {
		current_punc := ""
		if start(word, puncs) {
			current_punc = array[pos]
			if len(word) > 1 && !start(string(word[1]), puncs) {
				array[pos] = word[1:]
				array[pos-1] = array[pos-1] + word[:1]
			} else {
				array[pos-1] = array[pos-1] + current_punc
				array[pos] = ""
			}

		}
	}
	array = TrimEmpties(array)
	return array
}
func QuotesCheck(array []string) []string {
	puncs := []string{"'", "‘", "’", "\""}
	var checker bool
	var word_count int
	length := len(array) - 1
	//  Check word count between quotes
	for pos, word := range array {
		if strings.HasPrefix(array[0], "\"") {
			if pos != 0 {
				if strings.HasSuffix(array[pos-1], ":") && start(word, puncs) {
					checker = true
				}
			}
		}
		if checker && checkAlpha(word) {
			word_count++
		}
	}

	for pos, word := range array {
		var prev string
		var next string
		current := array[pos]
		// avoiding errors
		if pos != 0 && pos < length {
			prev = array[pos-1]
			next = array[pos+1]
		}
		if start(word, puncs) {
			if pos != 0 && pos < length {
				//  I ' m <-- case
				if checkAlpha(prev) && checkAlpha(next) {
					array[pos] = prev + current + next
				}
				//  Said: ' word <--case
				if strings.HasSuffix(prev, ":") {
					array[pos] += next
					array[pos+1] = ""
				}
			}
			//  if lastword is quote
			if start(array[length], puncs) && !checkAlpha(prev) {
				if strings.HasPrefix(array[length], "‘") {
					array[length] = "’"
				}
				array[length-1] += array[length]
				array[length] = ""
			}
		}
	}
	if word_count != 1 && checker {
		array[0] = array[0][1:]
		// checking word before last (last spot is "")
		array[length-1] = array[length-1][:len(array[length-1])-1]
	}
	array = TrimEmpties(array)
	return array
}
func CheckA(array []string) []string {
	for pos := range array {
		checks := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U", "h", "H"}
		var next string
		length := len(array) - 1
		// avoiding errors
		if pos != 0 && pos < length {
			next = array[pos+1]
		}
		if array[pos] == "a" && start(next, checks) {
			array[pos] = "an"
		}
	}
	TrimEmpties(array)
	return array
}
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {
	args := os.Args[1:]
	readable := args[0]
	if len(args) == 1 {
		fmt.Println("Too few arguments!")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("Too many arguments!")
		os.Exit(2)
	}
	file, err := os.Open(readable)
	CheckError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := ""
	content := ""
	// REading line by line
	for scanner.Scan() {
		line = scanner.Text()
		contentField1 := strings.Fields(string(line))
		contentField1 = BinToDec(contentField1)
		contentField1 = HexToDec(contentField1)
		contentField1 = ToUp(contentField1)
		contentField1 = ToLow(contentField1)
		contentField1 = Capitalize(contentField1)
		contentField1 = Ponctuation(contentField1)
		contentField1 = QuotesCheck(contentField1)
		contentField1 = CheckA(contentField1)
		line = strings.Join(contentField1, " ")
		content += line + "\n"
	}
	new_file, err := os.Create(args[1])
	CheckError(err)
	defer new_file.Close()
	new_file.WriteString(content)
}
