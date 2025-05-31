package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Walk scans the given path recursively up to maxDepth and prints a markdown-style tree.
// TODO: implement exclusion and output options.
func Walk(path string, maxDepth int) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	base := filepath.Base(path)
	suffix := ""
	if info.IsDir() {
		suffix = "/"
	}
	// print root node
	fmt.Printf("-   \n", base, suffix)
	if !info.IsDir() {
		return nil
	}
	return walk(path, 1, maxDepth)
}

// recursive helper to print directory tree
func walk(path string, depth, maxDepth int) error {
	if maxDepth > 0 && depth > maxDepth {
		return nil
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	// sort alphabetically
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})
	for _, e := range entries {
		name := e.Name()
		// skip hidden files/folders
		if strings.HasPrefix(name, ".") {
			continue
		}
		indent := strings.Repeat("  ", depth)
		display := name
		if e.IsDir() {
			display += "/"
		}
		fmt.Printf("%s-   %s\n", indent, display)
		if e.IsDir() {
			if err := walk(filepath.Join(path, name), depth+1, maxDepth); err != nil {
				return err
			}
		}
	}
	return nil
}
