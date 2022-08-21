package assembler

import (
	"testing"
)

func TestCoder_decimalToBin(t *testing.T) {
	type args struct {
		dec int
	}
	tests := []struct {
		name string
		c    *Coder
		args args
		want string
	}{
		{
			name: "十进制转二进制",
			c:    &Coder{},
			args: args{dec: 16},
			want: "000000000010000",
		},
		{
			name: "十进制转二进制",
			c:    &Coder{},
			args: args{dec: 0},
			want: "000000000000000",
		},
		{
			name: "十进制转二进制",
			c:    &Coder{},
			args: args{dec: 32768},
			want: "1000000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.decimalToBin(tt.args.dec); got != tt.want {
				t.Errorf("Coder.decimalToBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoder_stringToBinary(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		c    *Coder
		args args
		want int64
	}{
		{
			name: "",
			c:    &Coder{},
			args: args{str: "111111"},
			want: 0b111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.stringToBinary(tt.args.str); got != tt.want {
				t.Errorf("Coder.stringToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
