package main

import "testing"

type addTest struct {
	arg, expected string
}

var addTests = []addTest{
	addTest{"https://golang.org", "done"},
	addTest{"https://www.geeksforgeeks.org/", "done"},
	addTest{"https://www.javatpoint.com/java-tutorial", "done"},
	addTest{"https://www.w3schools.com/python/", "done"},
}

func TestFindImages1(t *testing.T) {
	for _, test := range addTests {
		if output := FindImages1(test.arg); output != test.expected {
			t.Errorf("failed")
		}
	}
}

func BenchmarkFindImages1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindImages1("https://golang.org")
	}
}

/*func TestFindImages1(t *testing.T) {
	//input := "https://golang.org"
	input := "https://www.geeksforgeeks.org/"
	expectedOutput := "done"
	output := FindImages1(input)

	if output != expectedOutput {
		t.Errorf("got %q, wanted %q", output, expectedOutput)
	}

}*/
