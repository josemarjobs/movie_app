package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var ActorNames []string

func Run(in stringReader, out io.Writer) {
	ActorNames = []string{}
	AskForNames(in)
	fmt.Fprintf(out, "You selected the following %d actors:\n", len(ActorNames))
	fmt.Fprintln(out, strings.Join(ActorNames, "\n"))
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
