package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Simple French-English dictionary
    dictionary := map[string]string{
        "chat":    "cat",
        "chien":   "dog",
        "maison":  "house",
        "pomme":   "apple",
        "livre":   "book",
        "soleil":  "sun",
        "lune":    "moon",
        "fleur":   "flower",
        "eau":     "water",
        "feu":     "fire",
    }

    fmt.Println("Mini French-English dictionary")
    fmt.Println("Type a French word (or 'exit' to leave):")
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        word, _ := reader.ReadString('\n')
        word = strings.TrimSpace(strings.ToLower(word))
        if word == "exit" {
            fmt.Println("Goodbye!")
            break
        }
        if translation, ok := dictionary[word]; ok {
            fmt.Printf("\"%s\" in English: %s\n", word, translation)
        } else {
            fmt.Println("Unknown word, try another one (or add it in the code!)")
        }
    }
}
