package hw_2

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestUnpackUsualCase(t *testing.T) {
	s := "a2b3cd4"
	actual, err := Unpack(s)

	require.Equal(t, "aabbbcdddd", actual, fmt.Sprintf("unpacking string: %s", s))
	require.Nil(t, err)
}

func TestUnpackSimplestCase(t *testing.T) {
	s := "abcd"
	actual, err := Unpack(s)

	require.Equal(t, "abcd", actual, fmt.Sprintf("unpacking string: %s", s))
	require.Nil(t, err)
}

func TestUnpackEmptyString(t *testing.T) {
	s := ""
	actual, err := Unpack(s)

	require.Equal(t, "", actual, fmt.Sprintf("unpacking string: %s", s))
	require.Nil(t, err)
}

func TestUnpackStringStartsWithDigit(t *testing.T) {
	s := "1abcd"
	actual, err := Unpack(s)

	require.Errorf(t, err, "string has to start from character, not a number", fmt.Sprintf("unpacking string: %s", s))
	require.Equal(t, "", actual)
}

func TestUnpackStringContainsNumbers(t *testing.T) {
	s := "a1bc48d"
	actual, err := Unpack(s)

	require.EqualError(t, err, "string can't contains numbers, you have to use only digits", fmt.Sprintf("unpacking string: %s", s))
	require.Equal(t, "", actual)
}

func BenchmarkUnpackUsualCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Unpack("a2b3cd4")
	}
}

func BenchmarkUnpackSimplestCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Unpack("abcd")
	}
}

func BenchmarkUnpackWithError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Unpack("abc48d")
	}
}

func BenchmarkUnpackEmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Unpack("")
	}
}