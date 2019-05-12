package problem0972

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	S   string
	T   string
	ans bool
}{

	{
		"1.0",
		"1",
		true,
	},

	{
		"350.(111)",
		"350.(11)",
		true,
	},

	{
		"15.(9)",
		"16",
		true,
	},

	{
		"1.0(9)",
		"1.1",
		true,
	},

	{
		"0.9(9)",
		"1.",
		true,
	},

	{
		"1.9(0)",
		"1.8(9)",
		true,
	},

	{
		"0",
		"0.",
		true,
	},

	{
		"0.(0)",
		"0",
		true,
	},

	{
		"0.(52)",
		"0.5(25)",
		true,
	},

	{
		"0.1666(6)",
		"0.166(66)",
		true,
	},

	// 可以有多个 testcase
}

func Test_isRationalEqual(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, isRationalEqual(tc.S, tc.T), "输入:%v", tc)
	}
}

func Benchmark_isRationalEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isRationalEqual(tc.S, tc.T)
		}
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		integer string
		nonRep  string
		repeat  string
	}{

		{
			"1",
			args{"1"},
			"1",
			"",
			"",
		},

		{
			"1.",
			args{"1."},
			"1",
			"",
			"",
		},

		{
			"1.0",
			args{"1.0"},
			"1",
			"",
			"",
		},

		{
			"1.0(9)",
			args{"1.0(9)"},
			"1",
			"0",
			"9",
		},

		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			integer, nonRep, repeat := parse(tt.args.s)
			if integer != tt.integer {
				t.Errorf("parse() integer = %v, want %v", integer, tt.integer)
			}
			if nonRep != tt.nonRep {
				t.Errorf("parse() nonRepeat = %v, want %v", nonRep, tt.nonRep)
			}
			if repeat != tt.repeat {
				t.Errorf("parse() repeat = %v, want %v", repeat, tt.repeat)
			}
		})
	}
}

func Test_normalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			"1",
			args{"1"},
			"10001",
		},

		{
			"1234.",
			args{"1234."},
			"11234",
		},

		{
			"1234.(1)",
			args{"1234.(1)"},
			"11234(1)",
		},

		// Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.s); got != tt.want {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
