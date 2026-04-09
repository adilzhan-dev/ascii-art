# ASCII Art

<https://01.tomorrow-school.ai/git/root/public/src/branch/master/subjects/ascii-art/README.md>

## Usage

```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
go run . "Hello World"| cat -e

go run . hello
go run . HELLO
go run . "{Hello & There #}"
go run . "1a\"#FdwHywR&/()="
go run . "{|}~"
go run . "[\]^_ 'a"
go run . ":;<=>?@"
go run . '\!" #$%&'"'"'()*+,-./'
go run "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run . abcdefghijklmnopqrstuvwxyz
```

Предусмотрен случай с текстом \n в качестве переноса строки вместо того чтобы передать более одного аргумента. Но вытекающие последсвия не исправлялись, такие как печать текста "\n" вместо переноса строки.

На этом этапе требуется только один шрифт (standard.txt). Символы не описаные в шаблоне (например кириллица) решено игнорировать.
