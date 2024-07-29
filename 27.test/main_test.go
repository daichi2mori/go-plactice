package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップさせる")
	}

	v := IsOne(i)
	// IsOneの返り値がtrueだとif文をスキップするためテストOK
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
