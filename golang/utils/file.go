package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func ReadFileIntoArrayOfInts(filename string) []int {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputAsStrings := strings.Split(string(content), "\n")
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
