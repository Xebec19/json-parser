package internal

import (
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []TokenList
		wantErr bool
	}{
		{
			name:  "single number",
			input: "5",
			want: []TokenList{
				{Key: SYMBOLS["NUMBER"], Val: "5"},
			},
		},
		{
			name:  "addition expression",
			input: "1+2",
			want: []TokenList{
				{Key: SYMBOLS["NUMBER"], Val: "1"},
				{Key: SYMBOLS["PLUS"], Val: "+"},
				{Key: SYMBOLS["NUMBER"], Val: "2"},
			},
		},
		{
			name:  "subtraction expression",
			input: "9-3",
			want: []TokenList{
				{Key: SYMBOLS["NUMBER"], Val: "9"},
				{Key: SYMBOLS["SUBST"], Val: "-"},
				{Key: SYMBOLS["NUMBER"], Val: "3"},
			},
		},
		{
			name:  "multiplication expression",
			input: "4*2",
			want: []TokenList{
				{Key: SYMBOLS["NUMBER"], Val: "4"},
				{Key: SYMBOLS["MULT"], Val: "*"},
				{Key: SYMBOLS["NUMBER"], Val: "2"},
			},
		},
		{
			name:  "division expression",
			input: "8/2",
			want: []TokenList{
				{Key: SYMBOLS["NUMBER"], Val: "8"},
				{Key: SYMBOLS["DIV"], Val: "/"},
				{Key: SYMBOLS["NUMBER"], Val: "2"},
			},
		},
		{
			name:  "expression with parentheses",
			input: "(1+2)",
			want: []TokenList{
				{Key: SYMBOLS["LPAREN"], Val: "("},
				{Key: SYMBOLS["NUMBER"], Val: "1"},
				{Key: SYMBOLS["PLUS"], Val: "+"},
				{Key: SYMBOLS["NUMBER"], Val: "2"},
				{Key: SYMBOLS["RPAREN"], Val: ")"},
			},
		},
		{
			name:  "complex expression",
			input: "(3+5)*2",
			want: []TokenList{
				{Key: SYMBOLS["LPAREN"], Val: "("},
				{Key: SYMBOLS["NUMBER"], Val: "3"},
				{Key: SYMBOLS["PLUS"], Val: "+"},
				{Key: SYMBOLS["NUMBER"], Val: "5"},
				{Key: SYMBOLS["RPAREN"], Val: ")"},
				{Key: SYMBOLS["MULT"], Val: "*"},
				{Key: SYMBOLS["NUMBER"], Val: "2"},
			},
		},
		{
			name:    "empty string",
			input:   "",
			want:    []TokenList{},
			wantErr: false,
		},
		{
			name:    "invalid character",
			input:   "1@2",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "letter character",
			input:   "a+1",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lexer(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Lexer(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("Lexer(%q) returned %d tokens, want %d", tt.input, len(got), len(tt.want))
				return
			}

			for i, token := range got {
				if token != tt.want[i] {
					t.Errorf("Lexer(%q) token[%d] = %+v, want %+v", tt.input, i, token, tt.want[i])
				}
			}
		})
	}
}
