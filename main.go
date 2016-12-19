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

	actors := []Actor{}
	var m sync.Mutex

	fmt.Fprintf(out, "\nYou selected the following %d actors:\n", len(ActorNames))

	var wg sync.WaitGroup
	for _, name := range ActorNames {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			actor, err := FetchActor(n)
			if err != nil {
				panic(err)
			}
			m.Lock()
			actors = append(actors, actor)
			m.Unlock()
			fmt.Fprintln(out, actor.Name)
		}(name)
	}
	wg.Wait()
	credits := FilterCredits(actors)
	if len(credits) > 0 {
		fmt.Fprintln(out, "They have appeared in the following movies and TV Shows together:")
		for _, c := range credits {
			fmt.Fprintln(out, c.ID, c.NameOrTitle())
		}
	} else {
		fmt.Fprintln(out, "\nThey haven't been in anything together.")
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
