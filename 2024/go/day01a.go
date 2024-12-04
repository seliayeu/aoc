package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "sort"
)

func main() {
    dat, err := os.ReadFile("day01.txt")

    if err != nil {
        panic(err)
    }

    lines := strings.Split(string(dat), "\n")
    numLines := len(lines) - 1
    
    l1 := make([]int, numLines)
    l2 := make([]int, numLines)

    for i, el := range(lines[:numLines]) {
        nums := strings.Split(el, " ")
        n1, err := strconv.Atoi(nums[0])
        if err != nil {
            panic(err)
        }
        n2, err := strconv.Atoi(nums[len(nums) - 1])
        if err != nil {
            panic(err)
        }
        l1[i] = n1
        l2[i] = n2
    }

    sort.Ints(l1)
    sort.Ints(l2)

    ans := 0

    for i := 0; i < numLines; i++ {
        ans += max(l1[i] - l2[i], -(l1[i] - l2[i]))
    }

    fmt.Print(ans)
}
