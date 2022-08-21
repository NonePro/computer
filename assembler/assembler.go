package assembler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var symbols = map[string]int{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"SCREEN": 16384,
	"KBD":    24576,
}

func main() {

}

type Assembler struct {
	p *Parser
	c *Coder
}

func NewAssembler(path string) *Assembler {
	return &Assembler{p: NewParser(path)}
}

func (a *Assembler) compile() {
	a.compileSymbol()
	a.p.reset()
	code := a.genBinCode("%04x")
	fmt.Print(code)
	binFile := "out.hack"
	file, err := os.OpenFile(binFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("v2.0 raw\n")
	writer.WriteString(code)
	writer.Flush()
}

func (a *Assembler) compileSymbol() {
	nextLine := 0
	for a.p.hasMoreLines() {
		a.p.next()
		fmt.Println(a.p.curInstruction, nextLine)
		if a.p.instructionType() == INSTRUCTION_L {
			symbol := a.p.symbol()
			if _, ok := symbols[symbol]; ok {
				panic(fmt.Sprintf("duplicate declaration %s", symbol))
			}
			symbols[symbol] = nextLine
		} else {
			nextLine++
		}
	}
}

func (a *Assembler) genBinCode(format string) string {
	code := ""
	variableAddr := 16
	for a.p.hasMoreLines() {
		a.p.next()
		switch a.p.instructionType() {
		case INSTRUCTION_A:
			addr, err := strconv.Atoi(a.p.symbol())
			if err != nil {
				// 变量声明
				addrNum, ok := symbols[a.p.symbol()]
				if !ok {
					addr = variableAddr
					symbols[a.p.symbol()] = addr
					variableAddr++
				} else {
					addr = addrNum
				}
			}
			addrCode := (a.c.decimalToBin(addr))[0:15]
			instruction := fmt.Sprintf("0%s", addrCode)
			fmt.Println(instruction)
			code += fmt.Sprintf(format, a.c.stringToBinary(instruction)) + "\n"
		case INSTRUCTION_C:
			destCode := a.c.destToBin(a.p.dest())
			compCode := a.c.compToBin(a.p.comp())
			jmpCode := a.c.jumpToBin(a.p.jump())
			instruction := fmt.Sprintf("111%s%s%s", compCode, destCode, jmpCode)
			fmt.Println(instruction)
			code += fmt.Sprintf(format, a.c.stringToBinary(instruction)) + "\n"
		}
	}
	return code
}
