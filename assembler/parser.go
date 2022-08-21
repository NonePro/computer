package assembler

import (
	"bufio"
	"os"
	"strings"
)

type Parser struct {
	lines          []string
	curInstruction Instruction
	curLine        int // 当前行号
}

func NewParser(path string) *Parser {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		t := scanner.Text()
		lines = append(lines, t)
	}

	return &Parser{
		lines: lines,
	}
}

func (p *Parser) hasMoreLines() bool {
	return p.curLine < len(p.lines)
}

func (p *Parser) reset() {
	p.curLine = 0
}

// 指针指到下条指令
func (p *Parser) next() {
	for {
		text := strings.Replace(p.lines[p.curLine], " ", "", -1)
		p.curLine++
		if string(text[0:2]) == "//" {
			continue
		}
		p.curInstruction = NewInstruction(text)
		break
	}
}

// 返回指令类型
func (p *Parser) instructionType() int {
	return p.curInstruction.Typ
}

// 返回符号
func (p *Parser) symbol() string {
	return p.curInstruction.Symbol
}

func (p *Parser) comp() string {
	return p.curInstruction.Comp
}

func (p *Parser) dest() string {
	return p.curInstruction.Dest
}

func (p *Parser) jump() string {
	return p.curInstruction.Jump
}

func (p *Parser) curLineNum() int {
	return p.curLine
}
