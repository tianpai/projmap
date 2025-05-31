package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tianpai/projmap/internal/tree"
)

func main() {
	// flags
	maxDepth := flag.Int("max-depth", 0, "limit recursion depth (0 = infinite)")
	exclude := flag.String("exclude", "", "Comma-separated list of patterns to ignore")
	// TODO: add --comment-map, --out, --plain, --show-hidden, --version
	flag.Parse()

	// target path (default = cwd)
	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	// build exclude list
	var excludes []string
	if *exclude != "" {
		for _, p := range strings.Split(*exclude, ",") {
			excludes = append(excludes, strings.TrimSpace(p))
		}
	}
	if err := tree.Walk(path, *maxDepth, excludes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
