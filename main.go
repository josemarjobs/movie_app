package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

var ActorNames []string

func Run(in stringReader, out io.Writer) {
	ActorNames = []string{}
	AskForNames(in)
	fmt.Fprintf(out, "You selected the following %d actors:\n", len(ActorNames))

	var wg sync.WaitGroup
	for _, name := range ActorNames {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			actor, err := FetchActor(n)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(out, actor.Name)
		}(name)
	}
	wg.Wait()
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
