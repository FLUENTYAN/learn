package chapter1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup2() {
	counts := make(map[string]int)
	file := os.Args[1]
	text, _ := os.Open(file)
	input := bufio.NewScanner(text)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, n := range counts {
		if n > 1 {
			fmt.Println("file")
		}
	}
}

func Dup3() {
	counts := make(map[string]int)
	file := os.Args[1]
	text, _ := ioutil.ReadFile(file)
	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		counts[line]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
