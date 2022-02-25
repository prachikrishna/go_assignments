package main

import "testing"

func TestFindImages1(t *testing.T) {
	//input := "https://golang.org"
	input := "https://www.geeksforgeeks.org/"
	expectedOutput := "done"
	output := FindImages1(input)

	if output != expectedOutput {
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}
