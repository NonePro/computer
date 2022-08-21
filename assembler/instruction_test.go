package assembler

import (
	"reflect"
	"testing"
)

func TestNewInstruction(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want Instruction
	}{
		{
			name: "A Instruction with constant",
			args: args{text: "@456"},
			want: Instruction{
				Typ:    INSTRUCTION_A,
				Comp:   "",
				Dest:   "",
				Jump:   "",
				Symbol: "456",
			},
		},
		{
			name: "A Instruction with variable",
			args: args{text: "@sum"},
			want: Instruction{
				Typ:    INSTRUCTION_A,
				Comp:   "",
				Dest:   "",
				Jump:   "",
				Symbol: "sum",
			},
		},
		{
			name: "L Instruction with variable",
			args: args{text: "#LOOP"},
			want: Instruction{
				Typ:    INSTRUCTION_L,
				Comp:   "",
				Dest:   "",
				Jump:   "",
				Symbol: "LOOP",
			},
		},
		{
			name: "C Instruction with comp and dest",
			args: args{text: "  M=D-M"},
			want: Instruction{
				Typ:    INSTRUCTION_C,
				Comp:   "D-M",
				Dest:   "M",
				Jump:   "",
				Symbol: "",
			},
		},
		{
			name: "C Instruction with comp and dest and space",
			args: args{text: "  M   =   D   -    M"},
			want: Instruction{
				Typ:    INSTRUCTION_C,
				Comp:   "D-M",
				Dest:   "M",
				Jump:   "",
				Symbol: "",
			},
		},
		{
			name: "C Instruction with comp and dest and jmp",
			args: args{text: "  M=D-M;JGT"},
			want: Instruction{
				Typ:    INSTRUCTION_C,
				Comp:   "D-M",
				Dest:   "M",
				Jump:   "JGT",
				Symbol: "",
			},
		},
		{
			name: "C Instruction with comp and dest and jmp",
			args: args{text: "  D-M;JGT"},
			want: Instruction{
				Typ:    INSTRUCTION_C,
				Comp:   "D-M",
				Jump:   "JGT",
				Symbol: "",
			},
		},
		{
			name: "C Instruction with comp and dest and jmp",
			args: args{text: "  D;JGT"},
			want: Instruction{
				Typ:    INSTRUCTION_C,
				Comp:   "D",
				Jump:   "JGT",
				Symbol: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInstruction(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
