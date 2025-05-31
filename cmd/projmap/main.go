package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tianpai/projmap/internal/tree"
)

func main() {
    // Define flags
    maxDepth := flag.Int("max-depth", 0, "Maximum depth to recurse (0 for unlimited)")
    flag.Parse()

    // Determine root path (positional argument)
    root := "."
    if args := flag.Args(); len(args) > 0 {
        root = args[0]
    }

    // Execute tree walk
    if err := tree.Walk(root, *maxDepth); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
