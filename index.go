package main

import (
	"fmt"
	"math/rand"
)

// type T int
// type K int

func find[T any](arr []T, cond func(T) bool) interface{} {
    for i := 0; i < len(arr); i++ {
        el := arr[i]
        if cond(el) {
            return el
        }
    }

    return nil
}

func main() {
    var arr = []int{1, 2, 3, 4, 5}

    randomIndex := rand.Intn(len(arr))
    fmt.Println("Random index:", randomIndex)

    x := find(arr, func(num int) bool {
        return num == arr[int(randomIndex)];
    })

    fmt.Println(x);
}