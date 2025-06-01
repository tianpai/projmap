package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tianpai/projmap/internal/tree"
)

// define GLOBAL variables for version and commit
const VERSION_NUM string = "0.0.1"

func main() {
	// normalize double-dash flags to single-dash for compatibility
	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "--") {
			os.Args[i] = "-" + strings.TrimPrefix(arg, "--")
		}
	}

	// flags
	maxDepth := flag.Int("max-depth", 0, "Limit recursion depth (0 = infinite)")
	exclude := flag.String("exclude", "", "Comma-separated list of patterns to exclude (e.g. *.tmp,*.log)")
	version := flag.Bool("version", false, "Display version and additional information")
	// TODO: add -comment-map, --out, --plain, --show-hidden
	flag.Parse()

	// NOTE: Final desicion:
	// only supports CWD where the command is issued
	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	// display all enabled arguments
	// BUG: this only works when one flag is used -> wrong print
	// fmt.Printf("Enabled flags: %v\n", os.Args[1:])

	// if version flag is set, print version and exit
	if *version {
		fmt.Printf("projmap version %s\n", VERSION_NUM)
		fmt.Println("License: Apache-2.0")
		fmt.Println("Issues: https://github.com/tianpai/projmap/issues")
		return
	}

	// build exclude list
	var excludes []string
	if *exclude != "" {
		// NOTE:
		// Ranging over SplitSeq is more efficient modernize (splitseq)
		for _, p := range strings.Split(*exclude, ",") {
			// fmt.Printf("Adding exclude pattern: %q\n", p)
			excludes = append(excludes, strings.TrimSpace(p))
		}
	}

	// TODO: add support for .gitignore files
	// by loading .gitignore into excludes if exists
	// MAKE SURE: .gitignore syntax is compatible with this tool
	fmt.Printf("Excludes: %v\n", excludes)
	if err := tree.Walk(path, *maxDepth, excludes); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
