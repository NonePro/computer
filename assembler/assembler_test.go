package assembler

import (
	"fmt"
	"testing"
)

func TestAssembler_compileSymbol(t *testing.T) {
	a := NewAssembler("sum1ToN.asm")
	a.compileSymbol()
	fmt.Println(symbols)
}

func TestAssembler_compile(t *testing.T) {
	a := NewAssembler("sum1ToN.asm")
	a.compile()
}
