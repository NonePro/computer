package assembler

import "strings"

const (
	INSTRUCTION_A = 1
	INSTRUCTION_C = 2
	INSTRUCTION_L = 3
)

type Instruction struct {
	Typ    int
	Comp   string
	Dest   string
	Jump   string
	Symbol string
}

func NewInstruction(text string) Instruction {
	text = strings.Replace(text, " ", "", -1)
	i := Instruction{}
	switch string(text[0]) {
	case "@":
		i.Typ = INSTRUCTION_A
		i.Symbol = text[1:]
	case "#":
		i.Typ = INSTRUCTION_L
		i.Symbol = text[1:]
	default:
		i.Typ = INSTRUCTION_C
		tokens := strings.Split(text, ";")
		if len(tokens) > 1 {
			i.Jump = strings.ToUpper(tokens[1])
		}
		tokens2 := strings.Split(tokens[0], "=")
		if len(tokens2) > 1 {
			i.Dest = tokens2[0]
			i.Comp = tokens2[1]
		} else {
			i.Comp = tokens2[0]
		}
	}
	return i
}
