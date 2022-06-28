package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString) FindVowels() []rune {
	var vowels []rune

	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func main() {
	name := MyString("Sam Aderson")
	fmt.Printf("Vowels are %c\n", name.FindVowels())
	var v VowelsFinder
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())
}
