package tree

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Walk scans the given path recursively up to maxDepth and prints a markdown-style tree.
// Excludes any file or folder whose name matches one of the patterns.
func Walk(path string, maxDepth int, excludes []string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// WARNING: the known issue
	// maybe becuase the path is set wrong here?
	base := filepath.Base(path)
	suffix := ""
	if info.IsDir() {
		suffix = "/"
	}
	// print root node
	fmt.Printf("-   %s%s\n", base, suffix)
	if !info.IsDir() {
		return nil
	}
	return walk(path, 1, maxDepth, excludes)
}

// recursive helper to print directory tree
func walk(path string, depth, maxDepth int, excludes []string) error {
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
		// skip excluded patterns
		skip := false
		for _, pat := range excludes {
			if pat == "" {
				continue
			}
			// DEBUG: Add logging
			// fmt.Printf("DEBUG: Checking %q against pattern %q\n", name, pat)
			// wildcard match
			// wildcard pattern: check for *, ?, [ or ]
			if strings.ContainsAny(pat, "*?[]") {
				if m, _ := filepath.Match(pat, name); m {
					// fmt.Printf("DEBUG: Wildcard match! Skipping %q\n", name)
					skip = true
					break
				}
			} else if name == pat || strings.Contains(name, pat) {
				// fmt.Printf("DEBUG: Exact/substring match! Skipping %q\n", name)
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		indent := strings.Repeat("  ", depth)
		display := name
		if e.IsDir() {
			display += "/"
		}
		fmt.Printf("%s-   %s\n", indent, display)
		if e.IsDir() {
			if err := walk(filepath.Join(path, name), depth+1, maxDepth, excludes); err != nil {
				return err
			}
		}
	}
	return nil
}
