package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	arraysstrings "github.com/rahul-as-dev/dsa-hub/arrays-strings"
)

func Solve() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	arr := buildArrayInput(reader, writer)
	fmt.Fprintln(writer, arraysstrings.NextGreaterElementToRight(arr))
}

func buildArrayInput(reader io.Reader, writer io.Writer) []int {
	var n int
	// It automatically skips spaces/newlines and converts types.
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &arr[i])
	}
	fmt.Fprintln(writer, arr)
	return arr
}