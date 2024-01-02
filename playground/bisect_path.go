package main

import (
    "fmt"
    "path"
    "strconv"
    "strings"
)

func bisectPath(p string) (head, tail string) {
    fmt.Println("\n" + "***** bisectPath *****")
    fmt.Println("p => " + p)

    p = path.Clean("/" + p)
    fmt.Println("path.Clean => " + p)

    i := strings.Index(p[1:], "/") + 1
    fmt.Println("p[1:] => " + p[1:])
    fmt.Println("i => " + strconv.Itoa(i))

    if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}

func main() {
    pathOne := "https://localhost:9000/places/123/"
    headOne, tailOne := bisectPath(pathOne)
    fmt.Println("***** pathOne *****")
    fmt.Println("pathOne => " + pathOne)
    fmt.Println("headOne => " + headOne)
    fmt.Println("tailOne => " + tailOne)

    pathTwo := "localhost:9000/places/123/"
    headTwo, tailTwo := bisectPath(pathTwo)
    fmt.Println("***** pathTwo *****")
    fmt.Println("pathTwo => " + pathTwo)
    fmt.Println("headTwo => " + headTwo)
    fmt.Println("tailTwo => " + tailTwo)

    pathThree := "/grandparent/parent/child/"
    headThree, tailThree := bisectPath(pathThree)
    fmt.Println("***** pathThree *****")
    fmt.Println("pathThree => " + pathThree)
    fmt.Println("headThree => " + headThree)
    fmt.Println("tailThree => " + tailThree)

    pathFour := "orphan"
    headFour, tailFour := bisectPath(pathFour)
    fmt.Println("***** pathFour *****")
    fmt.Println("pathFour => " + pathFour)
    fmt.Println("headFour => " + headFour)
    fmt.Println("tailFour => " + tailFour)
}

