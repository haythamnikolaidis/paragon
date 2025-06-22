package analysis

import (
	"fmt"
	"os"
	"strings"
)

func FindWordAtPosition(content string, line, character int) (string, error) {
    lines := strings.Split(content, "\n")
    if line < 0 || line >= len(lines) {
        return "", fmt.Errorf("line out of range")
    }

    words := lines[line]
    min := 0
    max := len(words)
    for i, c := range words {
        if c == ' ' {
            os.Stdout.WriteString(fmt.Sprintf("i: %d, min: %d, max: %d\n", i, min, max))
            if i > min && i < character {
                min = i
            }
            if i > character && i < max {
                max = i
                break;
            }
        }
    }

    return words[min:max], nil
}
