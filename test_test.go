package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkFast(b *testing.B) {
	Fast()
}
func BenchmarkSlow(b *testing.B) {
	Slow()
}
func Fast() {
	fmt.Printf(strings.Join(os.Args, " "))
}
func Slow() {
	var s, step string
	for _, arg := range os.Args {
		s += step + arg
		step = ""
	}
	fmt.Println(s)
}
