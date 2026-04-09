package converter

import (
	"os"
	"strings"
)

func loadChars(filename string) (map[rune][]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	result := make(map[rune][]string)
	lines := strings.Split(string(content), "\n")
	delta := int(' ')
	hasData := true
	for i := 0; hasData; i++ {
		r := rune((delta + i))
		result[r] = lines[i*(templateHeight)+1 : i*(templateHeight)+templateHeight]
		if r >= '~' {
			hasData = false
			break
		}
	}
	return result, nil
}

const (
	symbolHeight   = 8
	templateHeight = 9
)

var thinkerArt map[rune][]string

func init() {
	thinkerArt, _ = loadChars("./banners/standard.txt") //на этом этапе считаю, что ресурсы всегда доступны и в верном формате
}

// Преобразует текст в ASCII art
func ToAscii(text string) (string, error) {
	lines := strings.Split(text, "\n")
	result := ""
	for _, r := range lines {
		if r == "" {
			result += "\n"
			continue
		}

		for i := 0; i < symbolHeight; i++ {
			for _, c := range r {

				char, exists := thinkerArt[c]
				if !exists { //Отсутсвующие символы игнорировать
					continue
				}
				result += char[i]
			}
			result += "\n"
		}
	}
	return result, nil
}
