package utils

import (
	"bootstrap-application/start/alphabet"
)

func ConvertSentenceToArrayOfIndices(sentence string) []string {
	var sentenceToArrayIndices []string

	alphabetIndices := alphabet.GetAlphabetIndices()

	for _, letter := range sentence {
		letterString := string(letter)
		sentenceToArrayIndices = append(sentenceToArrayIndices, alphabetIndices[letterString])
	}

	return sentenceToArrayIndices
}

func SeparateArraysIntoFivePositions(array []string) [][]string {
	var separeArray [][]string

	rangeArray := len(array)
	indexMin := 0
	indexMax := 5

	for indexMax <= rangeArray {
		newArray := append(array[indexMin:indexMax])
		separeArray = append(separeArray, newArray)
		indexMin += 5
		indexMax += 5
	}

	return separeArray
}

func PanicConvertError() {
	panic("Could not convert text to integer.")
}
