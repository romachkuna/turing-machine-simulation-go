package main

// q0,[0,0,0,1,2,3],newState,[1,2,3,4,5,6],[L,R,R,R,R,R]
type TFunction struct {
	state      string
	tapeSymbol []string
	newSymbol  []string
	newState   string
	moves      []string
}

const BLANK = "#"
const L = "L"
const R = "R"
const N = "N"
const STARTSTATE = "q0"
