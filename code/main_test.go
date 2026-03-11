package main

import "testing"

func TestRunOnFile(t *testing.T) {
	status := processOneFile("../examples/ex12.bas", true)
	if status != 0 {
		t.Fail()
	}
}
