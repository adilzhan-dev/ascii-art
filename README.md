# ASCII Art

A command-line tool that converts text input into large graphic representations using ASCII characters (banner-style text art).

## 📖 Description

ASCII Art reads a string argument from the command line, maps each character to its multi-line ASCII banner equivalent (using the `standard` font), and prints the full output to the terminal. The program handles a wide range of input cases, including whitespace, punctuation, mixed case, and explicit newline sequences (`\n`) passed as literal text.

## 🎯 Key Features

- Converts arbitrary text into ASCII banner art
- Supports the full standard ASCII printable character range
- Handles multiple arguments and literal `\n` sequences as line breaks
- Gracefully ignores characters outside the supported font range (e.g. Cyrillic)

## 🛠️ Technologies

- Go (standard library only)

## 🚀 Usage

```bash
go run . "Hello World"
go run . "Hello\nThere"
go run . "1a\"#FdwHywR&/()="
```

Additional usage examples:

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "Hello\n\nThere" | cat -e
go run . hello
go run . HELLO
go run . "{Hello & There #}"
go run . "{|}~"
go run . "[\]^_ 'a"
go run . ":;<=>?@"
go run . abcdefghijklmnopqrstuvwxyz
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
```

## ⚙️ Implementation Notes

- Only one font is used at this stage (`standard.txt`)
- Characters not defined in the font template (e.g. Cyrillic letters) are intentionally ignored
- A special case is handled for the literal `\n` sequence passed within a single argument (as an alternative to passing multiple separate arguments) to trigger a line break. Some edge-case side effects of this approach (e.g. printing the literal text `\n` instead of a line break in specific combinations) were identified but left unresolved at this stage, as they fall outside the project's core requirements.

## 📂 Project Structure

```
ascii-art/
├── main.go
├── standard.txt      # font template
└── go.mod
```

## ⚠️ Known Limitations

Educational project built as part of the Tomorrow School (Astana Hub) curriculum — focused on file parsing, string/rune manipulation, and pattern matching in Go rather than production use.

## 👤 Authors

- Adilzhan Kusherbay ([@adilzhan-dev](https://github.com/adilzhan-dev))

## 📄 License

MIT
