package cryptografic

import (
	"bootstrap-application/start/alphabet"
	"bootstrap-application/start/utils"
	"math/rand"
	"strconv"
	"time"
)

var (
	Cryptograf string
	Key        []int
)

func Criptografic(setenceUpper string) {
	initialListWithIndexes := utils.ConvertSentenceToArrayOfIndices(setenceUpper)
	ListPopulatedWithIndexDivisibleByFive := popListWithNumberOfPositionsIsDivisibleByFive(initialListWithIndexes)
	arrayIntoFivePositions := utils.SeparateArraysIntoFivePositions(ListPopulatedWithIndexDivisibleByFive)
	Key = generateKeySecret(arrayIntoFivePositions)
	arrayIncrementPasskeyInTheIndexedPhrase := getIncrementPasskeyInTheIndexedPhrase(arrayIntoFivePositions, Key)
	Cryptograf = convertKeyedIndexToEncryptedText(arrayIncrementPasskeyInTheIndexedPhrase)
}

func popListWithNumberOfPositionsIsDivisibleByFive(array []string) []string {
	for len(array)%5 != 0 {
		array = append(array, "00")
	}

	return array
}



func generateKeySecret(array [][]string) []int {
	key := getInitialListForKeys()

	listWithTheHighestNumbersForEachPosition := getListWithTheHighestNumbersForEachPosition(array)

	for index, numberHighest := range listWithTheHighestNumbersForEachPosition {
		key[index] = 26 - numberHighest
	}

	return getRandomlyGeneratedNumberForEachSecretKeyIndex(key)
}

func getListWithTheHighestNumbersForEachPosition(array [][]string) []int {
	highNumberList := getInitialListForKeys()

	for _, listFivePositions := range array {

		for index, item := range listFivePositions {
			itemNumber, err := strconv.Atoi(item)

			if err != nil {
				utils.PanicConvertError()
			}
			highPositionNumber := highNumberList[index]

			if itemNumber > highPositionNumber {
				highNumberList[index] = itemNumber
			}
		}
	}

	return highNumberList
}

func getRandomlyGeneratedNumberForEachSecretKeyIndex(listMaximumNumberGeneratedRandomly []int) []int {
	key := getInitialListForKeys()

	// Generator random configuration.
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for index, item := range listMaximumNumberGeneratedRandomly {
		if item != 0 {
			random := r.Intn(item)

			if random == 0 {
				random += 1
			}

			key[index] = random

		}
	}

	return key
}

func getInitialListForKeys() []int {
	return []int{0, 0, 0, 0, 0}
}

func getIncrementPasskeyInTheIndexedPhrase(arrayIndixe [][]string, key []int) [][]string {
	var increment [][]string

	for _, fivePositionItem := range arrayIndixe {

		array := []string{}

		for index, item := range fivePositionItem {
			numberItem, err := strconv.Atoi(item)

			if err != nil {
				utils.PanicConvertError()
			}

			numberIndex := key[index]

			sum := numberIndex + numberItem
			var sumString string

			if sum <= 9 {
				sumString = "0" + strconv.Itoa(sum)
			} else {
				sumString = strconv.Itoa(sum)
			}

			array = append(array, sumString)
		}

		increment = append(increment, array)
	}

	return increment
}

func convertKeyedIndexToEncryptedText(arrayIndices [][]string) string {
	alphabetIndices := alphabet.GetAlphabetIndices()

	var criptograf string

	for _, array := range arrayIndices {

		for _, item := range array {
			criptograf += alphabet.GetKeyOfValue(alphabetIndices, item)
		}
	}

	return criptograf
}
