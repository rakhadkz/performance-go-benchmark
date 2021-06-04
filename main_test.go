package main

import (
	"bytes"
	"io/ioutil"
	"shyngys/my_solution"
	"testing"
)

func init() {
	OriginalSolution(ioutil.Discard)
	my_solution.MySolution(ioutil.Discard)
}

func TestMain(t *testing.T) {
	originalOut := new(bytes.Buffer)
	OriginalSolution(originalOut)
	originalResult := originalOut.String()

	fastOut := new(bytes.Buffer)
	my_solution.MySolution(fastOut)
	fastResult := fastOut.String()

	if originalResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, originalResult)
	}
}

func BenchmarkOriginalSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalSolution(ioutil.Discard)
	}
}

func BenchmarkMySolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		my_solution.MySolution(ioutil.Discard)
	}
}
