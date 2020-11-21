package hw_2

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestUnpackUsualCase(t *testing.T) {
	s := "a2b3cd4"
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, "aabbbcdddd", actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackSimplestCase(t *testing.T) {
	s := "abcd"
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, "abcd", actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackEmptyString(t *testing.T) {
	s := ""
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, "", actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackStringContainsEscapedDigits(t *testing.T) {
	s := `abcd\5`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, "abcd5", actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackUsualCaseContainsEscapedDigits(t *testing.T) {
	s := `\12a2b3\53c\3d4`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, "11aabbb555c3dddd", actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackUsualCaseContainsEscapedBackslash(t *testing.T) {
	s := `\\a2\\2b3\53c\3d4\\\`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, `\aa\\bbb555c3dddd\`, actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackTask1(t *testing.T) {
	s := `qwe\4\5`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, `qwe45`, actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackTask2(t *testing.T) {
	s := `qwe\45`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, `qwe44444`, actual, fmt.Sprintf("unpacking string: %s", s))
}

func TestUnpackTask3(t *testing.T) {
	s := `qwe\\5`
	actual, err := Unpack(s)

	require.Nil(t, err)
	require.Equal(t, `qwe\\\\\`, actual, fmt.Sprintf("unpacking string: %s", s))
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

func TestUnpackStringContainsNumbersAndEscapedBackslash(t *testing.T) {
	s := `a1bc\\48d`
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

func BenchmarkUnpackHeavyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Unpack(`\\a2\\2b3\53c\3d4\\\`)
	}
}
