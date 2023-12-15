package main 

import (
    "bufio"
    "fmt"
    "os"
)

type NumberTuple struct {
    Name string
    Number int
}

func sliceGen(names []string, numbers [] int) []NumberTuple {
    const iterLen int = 10
    result := make([] NumberTuple, 0)

    for i := 0; i <= iterLen; i++ {
        currentTuple := NumberTuple{
            Name: names[i],
            Number: numbers[i],
        }
        result = append(result, currentTuple)
    }
    return result
}

// func intPrint(typeSlice [][]interface{}) int {
//     for i := 0; i < len(typeSlice); i++ {
//         nameSlice := typeSlice[i][1]
//         numberSlice := typeSlice[i][2]
// 
//         currentSlice := sliceGen(nameSlice.([]string), numberSlice.([]int))
// 
//         fmt.Printf("\n%s\n\n", typeSlice[i][0])
// 
//         for j := 0; j < len(currentSlice); j++ {
//             fmt.Printf("%s::\t%d\n", currentSlice[j].Name, currentSlice[j].Number)
//         }
//     }
//     return 0
// }

func main() {
    decimalNames := []string{
        "zero",
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
        "ten",
    }

    decimalNumbers := []int{
        0,
        1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10,
    }

    binaryNames := []string{
        "0000",
        "0001",
        "0010",
        "0011",
        "0100",
        "0101",
        "0110",
        "0111",
        "1000",
        "1001",
        "1010",
    }

    binaryNumbers := []int{
        0b0,
        0b1,
        0b10,
        0b11,
        0b100,
        0b101,
        0b110,
        0b111,
        0b1000,
        0b1001,
        0b1010,
    }

    octalNames := []string{
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "10",
        "11",
        "12",
    }

    octalNumbers := []int{
        0o0,
        0o1,
        0o2,
        0o3,
        0o4,
        0o5,
        0o6,
        0o7,
        0o10,
        0o11,
        0o12,
    }

    hexadecimalNames := []string{
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        "a",
    }

    hexadecimalNumbers := []int{
        0x0,
        0x1,
        0x2,
        0x3,
        0x4,
        0x5,
        0x6,
        0x7,
        0x8,
        0x9,
        0xa,
    }

    decimalParams := []interface{}{
        "--Decimal--",
        decimalNames,
        decimalNumbers,
    }

    binaryParams := []interface{}{
        "--Binary--",
        binaryNames,
        binaryNumbers,
    }

    octalParams := []interface{}{
        "--Octal--",
        octalNames,
        octalNumbers,
    }

    hexadecimalParams := []interface{}{
        "--Hexadecimal--",
        hexadecimalNames,
        hexadecimalNumbers,
    }

    numericalTypes := [][]interface{}{
        decimalParams,
        binaryParams,
        octalParams,
        hexadecimalParams,
    }

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Printf("Input y to start\n\n")

    Loop:
    for scanner.Scan() {
        standardInput := scanner.Text()
        var continueToken bool

        if standardInput == "dec" {
            continueToken = true
        } else if standardInput == "bin" {
            continueToken = true
        } else if standardInput == "oct" {
            continueToken = true
        } else if standardInput == "hex" {
            continueToken = true
        } else {
            continueToken = false
        }

        if !continueToken {
            fmt.Println(standardInput)
            fmt.Printf("get outta here")
            break Loop
        }


        numericalType := 0

        switch {
        case standardInput == "dec":
            numericalType = 0
        case standardInput == "bin":
            numericalType = 1
        case standardInput == "oct":
            numericalType = 2
        case standardInput == "hex":
            numericalType = 3
    }

        nameSlice := numericalTypes[numericalType][1]
        numberSlice := numericalTypes[numericalType][2]



        currentSlice := sliceGen(nameSlice.([]string), numberSlice.([]int))

        fmt.Printf("\n%s\n\n", numericalTypes[numericalType][0])

        for j := 0; j < len(currentSlice); j++ {
            fmt.Printf("%s::\t%d\n", currentSlice[j].Name, currentSlice[j].Number)
        }

        fmt.Printf("Input y to continue\n\n")

    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}

