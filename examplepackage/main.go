package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:generate traceeverything

// Generate a random letter
func randomLetter() rune {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	return letters[rand.Intn(len(letters))]
}

// Generate a random word of a given length
func randomWord(length int) string {
	word := make([]rune, length)
	for i := range word {
		word[i] = randomLetter()
	}
	return string(word)
}

// Generate multiple random words
func randomWords(count, length int) []string {
	words := make([]string, count)
	for i := range words {
		words[i] = randomWord(length)
	}
	return words
}

// Generate a random sentence with a given number of words
func randomSentence(wordCount, wordLength int) string {
	words := randomWords(wordCount, wordLength)
	return strings.Join(words, " ") + "."
}

// Generate a random paragraph with a given number of sentences
func randomParagraph(sentenceCount, wordCount, wordLength int) string {
	sentences := make([]string, sentenceCount)
	for i := range sentences {
		sentences[i] = randomSentence(wordCount, wordLength)
	}
	return strings.Join(sentences, " ")
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	wordCount := 5
	wordLength := 8
	sentenceCount := 3

	words := randomWords(wordCount, wordLength)
	fmt.Println("Random Words:")
	for _, word := range words {
		fmt.Println(word)
	}

	sentence := randomSentence(wordCount, wordLength)
	fmt.Println("\nRandom Sentence:")
	fmt.Println(sentence)

	paragraph := randomParagraph(sentenceCount, wordCount, wordLength)
	fmt.Println("\nRandom Paragraph:")
	fmt.Println(paragraph)
}
