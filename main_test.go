package main

import (
	"testing"
)

func TestToUpperCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"hello", "HELLO"},
		{"Hello World", "HELLO WORLD"},
		{"", ""},
		{"already UPPER", "ALREADY UPPER"},
	}
	for _, tt := range tests {
		if got := toUpperCase(tt.input); got != tt.want {
			t.Errorf("toUpperCase(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"hello world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"hello_world", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"HelloWorld", "helloWorld"},
		{"one two three", "oneTwoThree"},
		{"", ""},
		{"single", "single"},
	}
	for _, tt := range tests {
		if got := toCamelCase(tt.input); got != tt.want {
			t.Errorf("toCamelCase(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"hello world", "hello_world"},
		{"Hello World", "hello_world"},
		{"helloWorld", "hello_world"},
		{"hello-world", "hello_world"},
		{"one two three", "one_two_three"},
		{"", ""},
		{"single", "single"},
	}
	for _, tt := range tests {
		if got := toSnakeCase(tt.input); got != tt.want {
			t.Errorf("toSnakeCase(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestToCapitalCase(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"hello world", "Hello World"},
		{"hello_world", "Hello World"},
		{"hello-world", "Hello World"},
		{"helloWorld", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"one two three", "One Two Three"},
		{"", ""},
		{"single", "Single"},
	}
	for _, tt := range tests {
		if got := toCapitalCase(tt.input); got != tt.want {
			t.Errorf("toCapitalCase(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

var benchInputs = []string{
	"hello",
	"hello world",
	"hello_world_foo_bar",
	"helloWorldFooBar",
	"Hello-World-Foo-Bar",
	"the quick brown fox jumps over the lazy dog",
}

func BenchmarkSplitWords(b *testing.B) {
	for _, input := range benchInputs {
		b.Run(input, func(b *testing.B) {
			for b.Loop() {
				splitWords(input)
			}
		})
	}
}

func BenchmarkToUpperCase(b *testing.B) {
	for _, input := range benchInputs {
		b.Run(input, func(b *testing.B) {
			for b.Loop() {
				toUpperCase(input)
			}
		})
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	for _, input := range benchInputs {
		b.Run(input, func(b *testing.B) {
			for b.Loop() {
				toCamelCase(input)
			}
		})
	}
}

func BenchmarkToSnakeCase(b *testing.B) {
	for _, input := range benchInputs {
		b.Run(input, func(b *testing.B) {
			for b.Loop() {
				toSnakeCase(input)
			}
		})
	}
}

func BenchmarkToCapitalCase(b *testing.B) {
	for _, input := range benchInputs {
		b.Run(input, func(b *testing.B) {
			for b.Loop() {
				toCapitalCase(input)
			}
		})
	}
}

func TestSplitWords(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"hello_world", []string{"hello", "world"}},
		{"hello-world", []string{"hello", "world"}},
		{"helloWorld", []string{"hello", "World"}},
		{"HelloWorld", []string{"Hello", "World"}},
		{"one_two-three four", []string{"one", "two", "three", "four"}},
	}
	for _, tt := range tests {
		got := splitWords(tt.input)
		if len(got) != len(tt.want) {
			t.Errorf("splitWords(%q) = %v, want %v", tt.input, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("splitWords(%q)[%d] = %q, want %q", tt.input, i, got[i], tt.want[i])
			}
		}
	}
}
