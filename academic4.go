package main

import(
    "fmt"
    "math/rand"
    "time"
)

func main() {
    head:= 0 
    tail := 0
    fmt.Println("Tossing a coin...")
    coin := make([]int, 3)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 3; i++ {
		coin[i] = rand.Intn(2)
        if coin[i] == 0 {
            fmt.Println("Round ", i+1, ":Heads")
            head++
        } else {
            fmt.Println("Round ", i+1, ":Tails")
            tail++
        }
    }
    fmt.Println("Heads:", head , ", Tails:", tail)
}


