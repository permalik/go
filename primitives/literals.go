package main 

import "fmt"

func main() {
    binaryZero := 0b0
    binaryOne := 0b1
    binaryTwo := 0b10
    binaryThree := 0b11
    binaryFour := 0b100
    binaryFive := 0b101
    binarySix := 0b110
    binarySeven := 0b111
    binaryEight := 0b1000
    binaryNine := 0b1001
    binaryTen := 0b1010

    octalZero := 0o0
    octalOne := 0o1
    octalTwo := 0o2
    octalThree := 0o3
    octalFour := 0o4
    octalFive := 0o5
    octalSix := 0o6
    octalSeven := 0o7
    octalEight := 0o10
    octalNine := 0o11
    octalTen := 0o12

    binarySlice := []int{
        binaryZero,
        binaryOne,
        binaryTwo,
        binaryThree,
        binaryFour,
        binaryFive,
        binarySix,
        binarySeven,
        binaryEight,
        binaryNine,
        binaryTen,
    }

    octalSlice := []int{
        octalZero,
        octalOne,
        octalTwo,
        octalThree,
        octalFour,
        octalFive,
        octalSix,
        octalSeven,
        octalEight,
        octalNine,
        octalTen,
    }

    counter := 0
    sliceLength := 10

    for ; counter <= sliceLength; {
        fmt.Printf(
            "binary: %v\toctal: %v\n",
            binarySlice[counter],
            octalSlice[counter],
            )
        counter++
    }
}

