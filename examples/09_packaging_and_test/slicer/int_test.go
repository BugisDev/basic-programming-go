package slicer_test

import (
	"testing"

	"github.com/bugisdev/basic-programming-go/examples/09_packaging_and_test/slicer"
)

func TestInt(t *testing.T) {
	var data slicer.Int
	data.Add(1)
	data.Add(2)

	if data.IndexOf(1) != 0 {
		t.Errorf("1 should be on index 0")
	}

	if data.Len() != 2 {
		t.Errorf("data should have the length of 2")
	}

	if data.IsExist(10) {
		t.Errorf("10 should not be exist")
	}

	var anotherdata = slicer.Int{3, 4}
	data.Merge(anotherdata)

	if data.Len() != 4 {
		t.Errorf("data should have the length of 4 now")
	}
}
