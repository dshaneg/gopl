package main

import (
	"fmt"
	"os"
	"strings"
)

type fileSet map[string]struct{}

func (fs fileSet) getKeys() []string {
	keys := make([]string, 0, len(fs))

	for k := range fs {
		keys = append(keys, k)
	}

	return keys
}

type dupEntry struct {
	count       int
	sourceFiles fileSet
}

func newDupEntry() *dupEntry {
	return &dupEntry{sourceFiles: make(fileSet)}
}

func (s *dupEntry) addDup(sourceFile string) {
	s.count++
	s.sourceFiles[sourceFile] = struct{}{}
}

func main() {
	counts := make(map[string]*dupEntry)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			set, ok := counts[line]
			if !ok {
				set = newDupEntry()
				counts[line] = set
			}
			set.addDup(filename)
		}
	}
	for line, set := range counts {
		if set.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", set.count, line, set.sourceFiles.getKeys())
		}
	}
}
