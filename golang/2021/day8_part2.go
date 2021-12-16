package main

import (
	"fmt"
	"github.com/adrian/adventofcode/utils"
	"os"
	"sort"
	"strings"
)

func findSignalsOfLength(signalPatterns []string, length int) (matchingSignalPatterns []string) {
	for _, pattern := range signalPatterns {
		if len(pattern) == length {
			matchingSignalPatterns = append(matchingSignalPatterns, pattern)
		}
	}
	return matchingSignalPatterns
}

func findUnknownLetter(pattern string, knownLetters []string) (unknownLetter string) {
	knownLettersString := strings.Join(knownLetters, "")
	for _, letter := range pattern {
		if !strings.Contains(knownLettersString, string(letter)) {
			unknownLetter = string(letter)
			break
		}
	}
	if unknownLetter == "" {
		fmt.Printf("Couldn't find unknown letter in [%s] with knownLetters: %v\n", pattern, knownLetters)
	}
	return unknownLetter
}

func decipherMappings(signalPatterns []string) (signalToSegmentMapping map[string]string) {
	signalCounts := make(map[string]int)
	for _, pattern := range signalPatterns {
		for _, char := range pattern {
			currentCount, _ := signalCounts[string(char)]
			signalCounts[string(char)] = currentCount + 1
		}
	}

	signalToSegmentMapping = make(map[string]string)
	for signal, count := range signalCounts {
		if count == 6 {
			signalToSegmentMapping["b"] = signal
		} else if count == 4 {
			signalToSegmentMapping["e"] = signal
		} else if count == 9 {
			signalToSegmentMapping["f"] = signal
		}
	}

	numberOneSignal := findSignalsOfLength(signalPatterns, 2)[0]
	unknownLetter := findUnknownLetter(numberOneSignal, []string{signalToSegmentMapping["f"]})
	signalToSegmentMapping["c"] = unknownLetter

	numberSevenSignal := findSignalsOfLength(signalPatterns, 3)[0]
	unknownLetter = findUnknownLetter(numberSevenSignal, []string{signalToSegmentMapping["f"], signalToSegmentMapping["c"]})
	signalToSegmentMapping["a"] = unknownLetter

	numberFourSignal := findSignalsOfLength(signalPatterns, 4)[0]
	unknownLetter = findUnknownLetter(numberFourSignal, []string{signalToSegmentMapping["f"], signalToSegmentMapping["c"], signalToSegmentMapping["b"]})
	signalToSegmentMapping["d"] = unknownLetter

	numberEightSignal := findSignalsOfLength(signalPatterns, 7)[0]
	unknownLetter = findUnknownLetter(numberEightSignal, []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["e"], signalToSegmentMapping["f"]})
	signalToSegmentMapping["g"] = unknownLetter

	return signalToSegmentMapping
}

func decipherNumber(signal string, signalToSegmentMapping map[string]string) (digit string) {
	zeroSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["c"], signalToSegmentMapping["e"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}
	oneSignals := []string{signalToSegmentMapping["c"], signalToSegmentMapping["f"]}
	twoSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["e"], signalToSegmentMapping["g"]}
	threeSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}
	fourSignals := []string{signalToSegmentMapping["b"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["f"]}
	fiveSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["d"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}
	sixSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["d"], signalToSegmentMapping["e"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}
	sevenSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["c"], signalToSegmentMapping["f"]}
	eightSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["e"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}
	nineSignals := []string{signalToSegmentMapping["a"], signalToSegmentMapping["b"], signalToSegmentMapping["c"], signalToSegmentMapping["d"], signalToSegmentMapping["f"], signalToSegmentMapping["g"]}

	sort.Strings(zeroSignals)
	sort.Strings(oneSignals)
	sort.Strings(twoSignals)
	sort.Strings(threeSignals)
	sort.Strings(fourSignals)
	sort.Strings(fiveSignals)
	sort.Strings(sixSignals)
	sort.Strings(sevenSignals)
	sort.Strings(eightSignals)
	sort.Strings(nineSignals)

	signalArray := strings.Split(signal, "")
	sort.Strings(signalArray)

	if strings.Join(signalArray, "") == strings.Join(zeroSignals, "") {
		digit = "0"
	} else if strings.Join(signalArray, "") == strings.Join(oneSignals, "") {
		digit = "1"
	} else if strings.Join(signalArray, "") == strings.Join(twoSignals, "") {
		digit = "2"
	} else if strings.Join(signalArray, "") == strings.Join(threeSignals, "") {
		digit = "3"
	} else if strings.Join(signalArray, "") == strings.Join(fourSignals, "") {
		digit = "4"
	} else if strings.Join(signalArray, "") == strings.Join(fiveSignals, "") {
		digit = "5"
	} else if strings.Join(signalArray, "") == strings.Join(sixSignals, "") {
		digit = "6"
	} else if strings.Join(signalArray, "") == strings.Join(sevenSignals, "") {
		digit = "7"
	} else if strings.Join(signalArray, "") == strings.Join(eightSignals, "") {
		digit = "8"
	} else if strings.Join(signalArray, "") == strings.Join(nineSignals, "") {
		digit = "9"
	} else {
		fmt.Println("Invalid signal pattern", signal)
		os.Exit(1)
	}

	return digit
}

func sumOutputValues(input string) int {
	signalPatterns := strings.Split(strings.TrimSpace(strings.Split(input, "|")[0]), " ")
	signalToSegmentMapping := decipherMappings(signalPatterns)

	valuePatterns := strings.Split(strings.TrimSpace(strings.Split(input, "|")[1]), " ")
	var numberString string
	for _, valuePattern := range valuePatterns {
		numberString = numberString + decipherNumber(valuePattern, signalToSegmentMapping)
	}

	return utils.Atoi(numberString)
}

func main() {
	text := utils.ReadFileIntoArrayOfStrings("day8_input")

	var totalOutputValues int
	for _, text := range text {
		if text == "" {
			continue
		}
		totalOutputValues = totalOutputValues + sumOutputValues(text)
	}

	fmt.Println(totalOutputValues)
}
