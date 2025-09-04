package main

import (
    "github.com/alhss/private-go-utils"
    "fmt"
)

func main() {
   version := utils.GetVersion()
   fmt.Printf("Using private-go-utils version: %s\n", version)
}
