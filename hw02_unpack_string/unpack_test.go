package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestOnlyNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "123123123", expected: ""},
		{input: "1", expected: ""},
		{input: "11111111", expected: ""},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSpecialChars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "~!&*()_+{}|:”>?<Ё!”№;%:?*()_+/Ъ,/.,;’[]\\|e", expected: "~!&*()_+{}|:”>?<Ё!”№;%:?*()_+/Ъ,/.,;’[]\\|e"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestBiggerThanByteChars(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "𒅑4𒀗3𒁅", expected: "𒅑𒅑𒅑𒅑𒀗𒀗𒀗𒁅"},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}
