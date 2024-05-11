package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readTransitionFunctions(filepath string) []TFunction {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var functions []TFunction

	for scanner.Scan() {
		line := scanner.Text()
		oldSymbols := line[strings.Index(line, "[") : strings.Index(line, "]")+2]
		line = strings.Replace(line, oldSymbols, "", 1)
		newSymbols := line[strings.Index(line, "[") : strings.Index(line, "]")+2]
		line = strings.ReplaceAll(line, newSymbols, "")
		moves := line[strings.Index(line, "[") : strings.Index(line, "]")+1]
		line = strings.ReplaceAll(line, moves, "")
		parts := strings.Split(line, ",")
		f := TFunction{
			state:      parts[0],
			tapeSymbol: strings.Split(oldSymbols[1:len(oldSymbols)-2], ","),
			newSymbol:  strings.Split(newSymbols[1:len(newSymbols)-2], ","),
			newState:   parts[1],
			moves:      strings.Split(moves[1:len(newSymbols)-2], ","),
		}
		functions = append(functions, f)
	}
	return functions
}

func readTape(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	tape := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		tape = append(tape, row)
	}

	return tape

}

func removeBlanks(s [][]string) [][]string {
	for i, e := range s {
		s[i] = remove(e)
	}
	return s
}

func remove(s []string) []string {
	out := make([]string, 0)
	for _, element := range s {
		if element != BLANK {
			out = append(out, element)
		}
	}
	return out
}
