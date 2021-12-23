package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoArrayOfStrings(filename string) (lines []string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines = strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
	}
	return lines
}

func ReadFileIntoArrayOfInts(filename string) []int {
	inputAsStrings := ReadFileIntoArrayOfStrings(filename)
	inputAsIntegers := []int{}
	for _, v := range inputAsStrings {
		if v == "" {
			continue
		}
		intVal, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		inputAsIntegers = append(inputAsIntegers, intVal)
	}
	return inputAsIntegers
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Problem opening file", filename, err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func ReadFileSingleLineOfInts(filename string) []int {
	contentsAsString := ReadFile(filename)[0]
	var intArray []int
	for _, s := range strings.Split(contentsAsString, ",") {
		intArray = append(intArray, Atoi(s))
	}
	return intArray
}
