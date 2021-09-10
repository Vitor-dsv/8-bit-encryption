package decryption

import (
	"bootstrap-application/start/alphabet"
	"bootstrap-application/start/utils"
	"strconv"
)

var (
	Decryption string
)

func Decrypt(setenceEncrypted string, key []int){
	arrayIndices := utils.ConvertSentenceToArrayOfIndices(setenceEncrypted)
	arrayIntoFivePositions := utils.SeparateArraysIntoFivePositions(arrayIndices)
	Decryption = addIndexesWithSecretKey(arrayIntoFivePositions, key)
}

func addIndexesWithSecretKey(array [][]string, key []int) string {
	var newArray string = ""

	alphabetIndices := alphabet.GetAlphabetIndices()

	for _, groupIndixes := range array {
		for indice, item := range groupIndixes {
			itemNumber, err := strconv.Atoi(item)
			if err != nil {
				utils.PanicConvertError()
			}

			sum := itemNumber - key[indice]
			sumString := strconv.Itoa(sum)
			if sum <= 9 {
				sumString = "0" + sumString
			}

			newArray += alphabet.GetKeyOfValue(alphabetIndices, sumString)
		}
	}

	return newArray
}

