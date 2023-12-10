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

    counter := 0

    for ; counter < len(binarySlice); {
        fmt.Println(binarySlice[counter])
        counter++
    }
}

