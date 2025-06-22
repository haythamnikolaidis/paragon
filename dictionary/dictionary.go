package dictionary

import (
	"fmt"
	"strings"

	"github.com/STRockefeller/dictionaries"
	"github.com/STRockefeller/dictionaries/instances/english/freedictionary"
)

func GetDefinition(word string) (string, error) {
	dictionary := dictionaries.NewEnglishDictionary()
	definition, err := dictionary.Search(word)
	if err != nil {
		return "", fmt.Errorf("failed to create Ollama client: %w", err)
	} else {
		result := FormatDefinition(definition[0])
		return result, nil
	}
}

func FormatDefinition(result freedictionary.ResultElement) string {
	var sb strings.Builder

	// Word and phonetic
	sb.WriteString("Word: " + result.Word)
	if result.Phonetic != "" {
		sb.WriteString("   [" + result.Phonetic + "]")
	}
	sb.WriteString("\n")

	// Origin
	if result.Origin != "" {
		sb.WriteString("Origin: " + result.Origin + "\n")
	}

	// Meanings
	for _, meaning := range result.Meanings {
		sb.WriteString("\n" + meaning.PartOfSpeech + "\n")
		for i, def := range meaning.Definitions {
			if i >= 2 {
				break
			}
			sb.WriteString(fmt.Sprintf("  %d. %s\n", i+1, def.Definition))
		}
	}

	return strings.TrimSpace(sb.String())
}
