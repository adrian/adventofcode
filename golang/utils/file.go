package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoArrayOfStrings(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(content), "\n")
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
