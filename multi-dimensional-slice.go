package main

import "fmt"

func main() {
    jb := []string{"James", "Bond", "action"}
    ts := []string{"Tony", "Stark", "science fiction"}
    
    xp := [][]string{jb, ts}
    
    fmt.Println(xp)
}
