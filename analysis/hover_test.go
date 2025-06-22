package analysis_test

import (
	"paragon/analysis"
	"testing"
)

func TestFindWordAtPosition(t *testing.T) {
    content := `
# Overview
The Paragon LSP is a Language Server Protocol (LSP) implementation in Go, designed to act as a writing assistant by integrating with the Ollama API. Its primary features include:
1. **Completions API**: Suggests words to complete a sentence.
2. **Hover API**: Displays the definition of a word when hovering over it.

This document provides a detailed and complete roadmap for implementing the Paragon LSP, including edge cases, concurrency handling, and compliance with the LSP specification.`
    line, character := 1, 3 
    expectedWord := "Overview"
    actualWord, err := analysis.FindWordAtPosition(content, line, character)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if expectedWord != actualWord {
        t.Fatalf("Expected: [%s], Actual: [%s]", expectedWord, actualWord)
    }
}
