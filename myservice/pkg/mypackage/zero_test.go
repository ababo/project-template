package mypackage_test

import (
	"testing"

	"pkg/mypackage"
)

func TestZero(t *testing.T) {
	if num := mypackage.Zero(); num != 0 {
		t.Errorf("Zero() expected to return 0, but returned %d", num)
	}
}
