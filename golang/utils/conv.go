package utils

import (
    "fmt"
    "math"
    "os"
    "strconv"
)

func ConvertBinaryStringToDecimal(number string) int {
    integer, err := strconv.Atoi(number)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
    return ConvertBinaryToDecimal(integer)
}

func ConvertBinaryToDecimal(number int) int {
    decimal := 0
    counter := 0.0
    remainder := 0

    for number != 0 {
        remainder = number % 10
        decimal += remainder * int(math.Pow(2.0, counter))
        number = number / 10
        counter++
    }
    return decimal
}
