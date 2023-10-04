package main

import "fmt"

func main() {
    var code string
    fmt.Scanln(&code)

    if len(code) != 2 {
        fmt.Println("Input non valido. Inserisci un codice a due cifre.")
        return
    }

    // Scambia le cifre
    swappedCode := string(code[1]) + string(code[0])

    fmt.Println(swappedCode)
}
