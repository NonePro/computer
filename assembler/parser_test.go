package assembler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		{
			name: "常规测试",
			args: args{path: "./a.txt"},
			want: &Parser{
				lines: []string{
					"hello world",
					"hack assembler by go",
					"",
					"java",
					"python",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitParser(t *testing.T) {
	p := NewParser("sum1ToN.asm")
	assert.Equal(t, len(p.lines), 28)
}

func TestNext(t *testing.T) {
	p := NewParser("sum1ToN.asm")
	for p.hasMoreLines() {
		p.next()
		str, err := json.Marshal(p.curInstruction)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s \n", str)
	}
}
