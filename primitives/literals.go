package main 

import "fmt"

type NumberTuple struct {
    Name string
    Number int
}

func sliceGen(names []string, numbers [] int) []NumberTuple{
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

    decimalParams := []interface{}{
        "decimal",
        decimalNames,
        decimalNumbers,
    }

    binaryParams := []interface{}{
        "binary",
        binaryNames,
        binaryNumbers,
    }

    octalParams := []interface{}{
        "octal",
        octalNames,
        octalNumbers,
    }

    numericalTypes := [][]interface{}{
        decimalParams,
        binaryParams,
        octalParams,
    }

    for i := 0; i < len(numericalTypes); i++ {
        nameSlice := numericalTypes[i][1]
        numberSlice := numericalTypes[i][2]
        
        currentSlice := sliceGen(nameSlice.([]string), numberSlice.([]int))

        fmt.Printf("\n%s\n\n", numericalTypes[i][0])

        for j := 0; j < len(currentSlice); j++ {
            fmt.Printf("%s::\t%d\n", currentSlice[j].Name, currentSlice[j].Number)
        }
    }
}

