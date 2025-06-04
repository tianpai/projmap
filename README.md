# projmap

## 📌 Summary

**Projmap** is a command-line tool that prints a markdown-friendly nested
file/folder structure with optional inline comments, similar to a tree command,
but designed for documentation, LLM inputs, and clean readability.

### 📄 Example Output

Running:

```bash
projmap -max-depth 3 -comment-map comments.yaml -exclude ".git,pycache"
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
projmap [options]
```

| Option          | Type   | Description                                       |
| --------------- | ------ | ------------------------------------------------- |
| `--max-depth`   | int    | Limit recursion depth                             |
| `--exclude`     | string | Comma-separated list of patterns to ignore        |
| `--comment-map` | string | Path to YAML/JSON file with comments per filename |
| `-out`          | string | Output to file instead of stdout                  |
| `-plain`        | flag   | Strip all comments                                |
| `--show-hidden` | flag   | Include dotfiles and hidden folders               |
| `-help`         | flag   | Show usage                                        |
| `-version`      | flag   | Show version                                      |

> [!WARNING]
> alias is not working yet, so use the full option names for now.

> [!TIP]
> The `--exclude` patterns support wildcards like `*.py` or `dir/*`
> to ignore directories use directory names WITHOUT trailing slashes.
> and directory names will be respected at all levels of the tree. (this behavior
> might change in the future)
> to ignore only files (aka. to show directory only )-> use `"*.*"`
> to ignore -> use `"*.*"`

## 📁 Comment Map Format

> [!warning]
> not supported yet

### YAML

```yaml
main.py: "Entry point"
helper.py: "Helper functions"
README.md: "Docs"
```

### JSON

```json
{
  "main.py": "Entry point",
  "helper.py": "Helper functions",
  "README.md": "Docs"
}
```

## 🚀 How to Run

Once the Go environment is ready, you can execute the CLI without installing:

```bash
# Run via go run
cd projmap
go run ./cmd/projmap/main.go [path] -max-depth=<n>

# e.g., scan current dir up to depth 3
go run ./cmd/projmap/main.go . -max-depth=3
```

## 🛠️ How to Compile

Build a standalone binary and run it:

```bash
# From project root
go build -o projmap ./cmd/projmap
# Verify and run
./projmap --max-depth=3
```

## 🚧 Development Status

### ✅ Completed Features

- **Core tree traversal** - Recursive directory parsing with depth control
- **Markdown output** - Clean `-` bullet format with 2-space indentation
- **Exclusion patterns** - Skip files/folders via `-exclude` flag (supports
  wildcards)
- **CLI interface** - Basic flag parsing with `-max-depth` and `-exclude`
- **Binary builds** - Compiles to standalone executable

### 🔄 In Progress

- **Comment maps** - YAML/JSON file annotation support (`-comment-map`)
- **Output redirection** - Write to file instead of stdout (`-out`)
- **Additional flags** - `-plain`, `-show-hidden`, `-version`

### 📋 TODO

- **Homebrew packaging** - Complete `.goreleaser.yml` and release automation
- **Testing** - Unit tests for walker and exclusion logic
- **Documentation** - Man page and usage examples
- **Performance** - Optimize for large directory trees

### 🐛 Known Issues

1. `--exclude` patterns do not work when given a PATH
   `projmap -max-depth 3 -exclude ".git,pycache"` works, but
   `projmap . -max-depth 3 -exclude ".git,pycache"` does not

_Last updated: May 31, 2025_
