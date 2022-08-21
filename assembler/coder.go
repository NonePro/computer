package assembler

import (
	"fmt"
	"strconv"
)

var destMap = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"A":   "100",
	"AM":  "101",
	"AD":  "110",
	"ADM": "111",
}

var compMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

var jumpMap = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

type Coder struct {
}

func (c *Coder) decimalToBin(dec int) string {
	return fmt.Sprintf("%015b", dec)
}

func (c *Coder) destToBin(dest string) string {
	code, ok := destMap[dest]
	if !ok {
		panic("undefined dest")
	}
	return code
}

func (c *Coder) compToBin(comp string) string {
	code, ok := compMap[comp]
	if !ok {
		panic(fmt.Sprintf("undefined comp %s", comp))
	}
	return code
}

func (c *Coder) jumpToBin(jmp string) string {
	code, ok := jumpMap[jmp]
	if !ok {
		panic("undefined jmp")
	}
	return code
}

func (c *Coder) stringToBinary(str string) int64 {
	s, err := strconv.ParseInt(str, 2, 32)
	if err != nil {
		panic(err)
	}
	return s
}
