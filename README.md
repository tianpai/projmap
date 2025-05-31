# projmap

### 📌 Summary

projmap is a command-line tool that prints a markdown-friendly nested
file/folder structure with optional inline comments, similar to a tree command,
but designed for documentation, LLM inputs, and clean readability.

## 📄 Example Output

Running:

```bash
projmap . --max-depth 3 --comment-map comments.yaml --exclude ".git,pycache"
```

Output:

```txt
-   project-root/
-   src/
    -   main.py # Entry point
    -   utils/
        -   helper.py # Helper functions
-   tests/
    -   test_main.py # Unit tests
-   requirements.txt # Dependencies
-   README.md # Docs
```

## ⚙️ Features

### 🔸 Core

| Feature               | Description                                                     |
| --------------------- | --------------------------------------------------------------- |
| Markdown-style output | Nested list using `-` and 2-space indents, folders end with `/` |
| Inline comments       | Auto-generated or loaded from YAML/JSON via `--comment-map`     |
| File/folder exclusion | Use `--exclude "pattern1,pattern2"` to skip files/folders       |
| Depth control         | `--max-depth N` to limit how deep to recurse                    |
| Comment map support   | `--comment-map path.yaml` to annotate files (optional)          |
| Output file option    | `--out file.md` to write output to file                         |
| Ignore hidden files   | By default, skip dotfiles unless `--show-hidden` is passed      |
| Sorted output         | Alphabetical within folders                                     |

### 🔸 Command Line Options

```bash
projmap [path] [options]
```

| Option          | Type   | Description                                       |
| --------------- | ------ | ------------------------------------------------- |
| `--max-depth`   | int    | Limit recursion depth                             |
| `--exclude`     | string | Comma-separated list of patterns to ignore        |
| `--comment-map` | string | Path to YAML/JSON file with comments per filename |
| `--out`         | string | Output to file instead of stdout                  |
| `--plain`       | flag   | Strip all comments                                |
| `--show-hidden` | flag   | Include dotfiles and hidden folders               |
| `--help`        | flag   | Show usage                                        |
| `--version`     | flag   | Show version                                      |

## 📁 Comment Map Format

### YAML

main.py: "Entry point" helper.py: "Helper functions" README.md: "Docs"

### JSON

{ "main.py": "Entry point", "helper.py": "Helper functions", "README.md": "Docs"
}

## 📦 Homebrew-Ready Design

-   Written in **Go** (static binary)
-   GitHub-based releases with version tags (v1.0.0)
-   Use goreleaser to automate:
-   Binary builds
-   Checksums
-   Homebrew formula
-   Optional: Your own tap repo (e.g. yourname/homebrew-projmap)

## 🧱 Directory Structure (Go Project)

```
projmap/
├── cmd/
│   └── projmap/
│       └── main.go
├── internal/
│   └── tree/
│       └── walker.go
├── go.mod
├── .goreleaser.yml
├── README.md
└── LICENSE
```

## ✅ MVP Checklist

-   Parses a directory recursively
-   Outputs markdown-style tree
-   Respects --max-depth
-   Skips excluded folders/files
-   Adds inline comments from optional YAML/JSON
-   Writes to file (or stdout)
-   CLI interface
-   Buildable Go binary
-   Homebrew-compatible release setup

## 🚀 How to Run

Once the Go environment is ready, you can execute the CLI without installing:

```bash
# Run via go run
cd projmap
go run ./cmd/projmap/main.go [path] --max-depth=<n>
# e.g., scan current dir up to depth 3
go run ./cmd/projmap/main.go . --max-depth=3
```

## 🛠️ How to Compile

Build a standalone binary and run it:

```bash
# From project root
go build -o projmap ./cmd/projmap
# Verify and run
./projmap [path] --max-depth=3
```

### 📥 Install as a global command

To install into your $GOBIN (or $GOPATH/bin) so it's on your PATH:

```bash
go install github.com/tianpai/projmap/cmd/projmap@latest
# Then invoke anywhere:
projmap . --max-depth=3
```
