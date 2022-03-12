package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	err := Copy("1.txt", "2.txt", 0, 0)
	require.Error(t, err)

	err = Copy("testdata/input.txt", "2.txt", 10000, 0)
	require.Error(t, err)
	require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "error %q", err)

	err = Copy("/dev/urandom", "urandom", 0, 0)
	require.Error(t, err)
	require.Truef(t, errors.Is(err, ErrUnsupportedFile), "error %q", err)
}
