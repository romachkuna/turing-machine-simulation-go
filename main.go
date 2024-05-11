package main

import "fmt"

func main() {
	tape := readTape("tape.txt")
	instructions := readTransitionFunctions("instructions.txt")
	result := startTuringMachine(tape, instructions)
	fmt.Println(result)
}

func startTuringMachine(tape [][]string, tFunctions []TFunction) [][]string {
	tapeSize := len(tape)
	currentState := STARTSTATE
	done := false
	heads := make(map[int]int)

	for i, _ := range tape {
		heads[i] = 0
	}

	for !done {
		for _, t := range tFunctions {
			flag := false
			if t.state == currentState {
				// checking for right transition function
				for i := 0; i < tapeSize; i++ {
					if tape[i][heads[i]] == t.tapeSymbol[i] {
						if i == tapeSize-1 {
							flag = true
						}
					} else {
						break
					}
				}
				// correct transition function found
				if flag {
					// update the tape
					for i, _ := range tape {
						tape[i][heads[i]] = t.newSymbol[i]
					}
					// update state
					currentState = t.newState

					// update heads
					for i, m := range t.moves {
						switch m {
						case N:
						case R:
							heads[i] = heads[i] + 1
							// append blank to the tape
							if heads[i] >= len(tape[i]) {
								tape[i] = append(tape[i], BLANK)
							}
						case L:
							if heads[i] <= 0 {
								// prepend blank to the tape
								tape[i] = append([]string{BLANK}, tape[i]...)
							}
							heads[i] = heads[i] - 1
						}
					}
					break
				}
			}
		}
		if currentState == "halt" {
			done = true
		}
	}
	return removeBlanks(tape)

}
