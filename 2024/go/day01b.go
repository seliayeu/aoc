package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    // "sort"
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
    m := make(map[int]int)


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
        
        if val, ok := m[n2]; ok {
            m[n2] = val + 1
        } else {
            m[n2] = 1
        }

    }

    // sort.Ints(l1)
    // sort.Ints(l2)
    //
    ans := 0

    for i := 0; i < numLines; i++ {
        if val, ok := m[l1[i]]; ok {
            ans += l1[i] * val
        }
    }

    fmt.Print(ans)
}
