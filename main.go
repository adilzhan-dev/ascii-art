package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"ascii-art/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <text>")
		return
	}
	text := strings.ReplaceAll(os.Args[1], "\\n", "\n") //На всякий случай. Для аудиторов вдруг кто-то буквально будет использовать \n как Enter. Во всяком случае в услвии задачи такой сценарий прописан - его и придерживаемся. Варианты типа \\n обрабатывать не будем без явного запроса.
	spaceInput := regexp.MustCompile(`^\s*$`)           //Тоже доп условие из задания. На всякий случай делаем как в примере.
	if spaceInput.MatchString(text) {
		fmt.Print(text)
		return
	}
	result, err := converter.ToAscii(text)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = strings.ReplaceAll(result, "\n\n\n\n\n\n\n\n", "\n") //Не хочу менять основной алгоритм (думаю в будущем пригодится основной сценарий), но в задаче есть примеры, где "Большой Enter" превращается в "Маленький Enter". Вот и убираем лишние пустые строки.
	fmt.Print(result)
}
