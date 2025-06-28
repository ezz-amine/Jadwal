package core

import (
	"strings"
)

func CenterPad(text string, char rune, totalLength int) string {
	if len(text) >= totalLength {
		return text
	}
	padding := (totalLength - len(text)) / 2
	return strings.Repeat(string(char), padding) + text +
		strings.Repeat(string(char), totalLength-len(text)-padding)
}

func splitLongWord(word string, leadOffset int, chunkSize int) []string {
	const startOfChunk = "~"
	const endOfChunk = "$"
	const ellipsis = "..."

	if len(word) < 1 {
		return []string{word}
	}
	if len(word) < chunkSize {
		return []string{word}
	}

	if chunkSize < 1 {
		chunkSize = 1 + len(endOfChunk)
	}
	if leadOffset < 0 || leadOffset > chunkSize {
		leadOffset = 0
	}

	start := 0
	end := chunkSize - leadOffset - 1
	firstPart := word[start:end] + endOfChunk
	secondPart := startOfChunk + word[end:]

	if len(secondPart) > chunkSize {
		secondPart = secondPart[:chunkSize-len(ellipsis)] + ellipsis
	}

	return []string{firstPart, secondPart}
}

func appendWord(currentChunk string, word string, maxLength int) []string {
	currentSize := len(currentChunk)
	wordLenght := len(word)
	nextSize := currentSize + wordLenght
	if nextSize+1 < maxLength {
		return []string{currentChunk + word + " "}
	}

	if nextSize+1 == maxLength {
		return []string{currentChunk + word, ""}
	}

	if wordLenght > maxLength {
		wordSegments := splitLongWord(word, currentSize, maxLength)
		if len(wordSegments) == 2 {
			wordSegments[0] = currentChunk + wordSegments[0]
			if len(wordSegments) == maxLength {
				return append(wordSegments, "")
			}
			return wordSegments
		} else {
			word = wordSegments[0]
		}
	}

	return []string{strings.TrimSpace(currentChunk), word + " "}
}

func SplitTextByWords(text string, maxLength int) []string {
	if maxLength <= 0 {
		return []string{text}
	}

	words := strings.Fields(text) // Split into words
	if len(words) == 0 {
		return []string{}
	}

	var chunks []string
	currentChunk := words[0] + " "
	for _, word := range words[1:] {
		newChunks := appendWord(currentChunk, word, maxLength)
		currentChunk = newChunks[len(newChunks)-1]
		chunks = append(chunks, newChunks[:len(newChunks)-1]...)
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
